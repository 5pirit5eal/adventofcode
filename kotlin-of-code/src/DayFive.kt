import io.github.oshai.kotlinlogging.KotlinLogging
import java.io.File
import java.rmi.UnexpectedException

private val logger = KotlinLogging.logger {}

fun fifthDay(): Pair<Int, Int> {
    val rules = HashMap<String, HashMap<String, MutableList<String>>>()
    val manuals: MutableList<List<String>> = mutableListOf()
    val filePath = "inputs/fifth_day_input.txt"
    // read txt file and construct hashmap and lists
    val lines = File(filePath).useLines { lines ->
        lines.forEach { line ->
            if ("|" in line) {
                // construct hashmap
                val (firstPage, secondPage) = line.split("|")
                if (rules.containsKey(firstPage)) {
                    rules[firstPage]!!["infront"]!!.add(secondPage)
                } else {
                    rules[firstPage] =
                        hashMapOf("infront" to mutableListOf<String>(secondPage), "after" to mutableListOf<String>())
                }
                if (rules.containsKey(secondPage)) {
                    rules[secondPage]!!["after"]!!.add(firstPage)
                } else {
                    rules[secondPage] =
                        hashMapOf("after" to mutableListOf<String>(firstPage), "infront" to mutableListOf<String>())
                }
            } else if ("," in line) {
                manuals.add(line.split(","))
            }
        }
    }
    val rulesInput = rules.mapValues { firstLevel -> firstLevel.value.mapValues { it.value.toList() } }

    val outputs = mutableListOf<String>()
    val orderedOutputs = mutableListOf<String>()
    for (manual in manuals) {
        if (checkOrder(rulesInput, manual)) {
            outputs.add(manual[manual.size / 2])
        } else {
            val orderedManual = reorderManual(rulesInput, manual)
            if (checkOrder(rulesInput, orderedManual)) {
                orderedOutputs.add(orderedManual[orderedManual.size / 2])
            } else {
                throw UnexpectedException("Ordering didn't work for $manual")
            }
        }
    }
    return outputs.sumOf { it.toInt() } to orderedOutputs.sumOf { it.toInt() }
}

fun checkOrder(rules: Map<String, Map<String, List<String>>>, manual: List<String>): Boolean {
    // Checks if the given manual is in correct order according to rules
    // Iterate over the manual and look back and ahead if everything fits
    for ((i, page) in manual.withIndex()) {
        if (!rules.containsKey(page)) {
            continue
        }
        val followingPages = rules[page]!!["infront"]!!
        val previousPages = rules[page]!!["after"]!!
        for (n in 0..<i) {
            if (manual[n] in followingPages) {
                logger.info { "For $page at $i: ${manual[n]} at $n is not allowed after according to ${rules[page]!!}" }
                return false
            }
        }
        for (m in i + 1..<manual.size) {
            if (manual[m] in previousPages) {
                logger.info { "For $page at $i: ${manual[m]} at $m is not allowed before according to ${rules[page]!!}" }
                return false
            }
        }
    }
    return true
}

fun reorderManual(rules: Map<String, Map<String, List<String>>>, manual: List<String>): List<String> {
    return manual.sortedWith(Comparator { a, b ->
        when {
            rules[a]?.get("infront")?.contains(b) == true -> -1
            rules[a]?.get("after")?.contains(b) == true -> 1
            else -> 0
        }
    })
}