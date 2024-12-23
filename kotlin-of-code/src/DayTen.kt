import io.github.oshai.kotlinlogging.KotlinLogging
import java.io.File
import kotlin.math.max
import kotlin.math.min

private val logger = KotlinLogging.logger {}

fun tenthDay(): Int {
    val filePath = "inputs/tenth_day_input.txt"
    // read txt file and construct 2D matrix
    val lines = File(filePath).readLines()
    val matrix = Array(lines.size) { IntArray(lines[0].length) }
    val zeroPositions = mutableSetOf<Pair<Int, Int>>()

    for (i in lines.indices) {
        for (j in lines[i].indices) {
            matrix[i][j] = lines[i][j].digitToInt()
            if (matrix[i][j] == 0) {
                zeroPositions.add(i to j)
            }
        }
    }
    var result = 0
    for (zeroPosition in zeroPositions) {
        result += findPaths(matrix, zeroPosition)
    }

    return result
}

fun findPaths(matrix: Array<IntArray>, startPosition: Pair<Int, Int>): Int {
    val pointQueue = ArrayDeque<Triple<Int, Int, Int>>(listOf(Triple(startPosition.first, startPosition.second, 0)))
    val ninePositions = mutableSetOf<Pair<Int, Int>>()
    val visitedPositions = HashMap<String, MutableSet<Pair<Int, Int>>>()
    val xDim = matrix.size - 1
    val yDim = matrix[0].size - 1
    while (pointQueue.isNotEmpty()) {
        val currentPoint = pointQueue.removeLast()

        // construct cross
        val cross = listOf(
            Pair(max(currentPoint.first - 1, 0), currentPoint.second),
            Pair(min(currentPoint.first + 1, xDim), currentPoint.second),
            Pair(currentPoint.first, max(currentPoint.second - 1, 0)),
            Pair(currentPoint.first, min(currentPoint.second + 1, yDim))
        )
        for ((i, j) in cross) {
            val matrixValue = matrix[i][j]
            if (matrixValue == currentPoint.third + 1 && matrixValue == 9) {
                ninePositions.add(i to j)
                continue
            }
            if (matrixValue == currentPoint.third + 1) {
                pointQueue.addFirst(Triple(i, j, matrixValue))
            }
        }
    }

    return ninePositions.count()
}