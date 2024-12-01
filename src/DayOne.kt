import java.io.File
import kotlin.math.abs

fun parseInput(): Pair<Array<Int>, Array<Int>> {
    val filePath = "inputs/first_day_input.txt"
    val a = mutableListOf<Int>()
    val b = mutableListOf<Int>()
    // read txt file
    // iterate and append to output lists
    File(filePath).useLines { lines ->
        lines.forEach {
            val lineList = it.split(" ")
            a.add(lineList.first().toInt())
            b.add(lineList.last().toInt())
        }
    }
    return Pair(a.toTypedArray(), b.toTypedArray())
}


fun measureDistance(a: Array<Int>, b: Array<Int>): Int {
    // sort the arrays
    a.sort()
    b.sort()
    if (a.contentEquals(b)) {
        return 0
    }

    // loop over array and calculate distance
    var distance: Int = 0
    for (i in a.indices) {
        distance += abs(a[i] - b[i])
    }
    // return distance
    return distance
}

fun measureSimilarity(a: Array<Int>, b: Array<Int>): Int {
    val occurrencesPerNumber = HashMap<Int, Int>()

    for (num in b) {
        occurrencesPerNumber[num] = (occurrencesPerNumber[num] ?: 0) + 1
    }

    // loop over array and calculate distance
    // var similarity: Int = 0
    // for (i in a) {
    //     similarity += i * (occurrencesPerNumber[i] ?: 0) for i in a
    // }
    // return distance
    return a.sumOf { num -> num * (occurrencesPerNumber[num] ?: 0)}
}

fun measureSimilaritySolution(a: Array<Int>, b: Array<Int>): Int {
    val occurrencesPerNumber = b.groupingBy { it }.eachCount()
    return a.sumOf { num -> num * (occurrencesPerNumber[num] ?: 0) }
}