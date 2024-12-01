import java.io.File
import kotlin.math.abs

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
fun main() {
    val (a, b) = parseInput()
    println(measureDistance(a, b))
}

fun parseInput(): Pair<Array<Int>, Array<Int>> {
    val filePath = "inputs/first_day_input.txt"
    var a = mutableListOf<Int>()
    var b = mutableListOf<Int>()
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
