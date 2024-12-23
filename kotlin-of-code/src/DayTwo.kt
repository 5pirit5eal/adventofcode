import java.io.File
import kotlin.math.abs

fun secondDay(): Int {
    val filePath = "inputs/second_day_input.txt"
    var numSafe = 0
    // read txt file
    // iterate and append to output lists
    File(filePath).useLines { lines ->
        lines.forEach {
            val lineList = it.split(" ").map { num -> num.toInt() }
            if (problemDampener(lineList)) {
                numSafe++
                // println(lineList)
            }
        }
    }
    return numSafe
}


fun safetyCheck(report: List<Int>): Boolean {
    val ascending = report[0] < report[1]
    for (idx in 1..<report.size) {
        val diffLower = report[idx - 1] - report[idx]
        if (diffLower == 0 || abs(diffLower) !in 1..3) {
            return false
        }
        if (((diffLower < 0) && !ascending) || ((diffLower > 0) && ascending)) {
            return false
        }
    }
    return true
}

fun problemDampener(report: List<Int>): Boolean {
    val originalReport = report

    for (idx in originalReport.indices) {
        var reducedReport = report.toMutableList()
        reducedReport.removeAt(idx)
        if (safetyCheck(reducedReport)) {
            return true
        }
    }

    return false
}