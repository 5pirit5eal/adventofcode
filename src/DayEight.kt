import io.github.oshai.kotlinlogging.KotlinLogging
import java.io.File

private val logger = KotlinLogging.logger {}

fun eightDay(): Int {
    val filePath = "inputs/eight_day_input.txt"
    // read txt file and construct 2D matrix
    val lines = File(filePath).readLines()
    val matrix = Array(lines.size) { CharArray(lines[0].length) }

    for (i in lines.indices) {
        for (j in lines[i].indices) {
            matrix[i][j] = lines[i][j]
        }
    }
    return getAntinodes(matrix)
}

private fun getAntinodes(matrix: Array<CharArray>): Int {
    val antinodePositions = mutableSetOf<Pair<Int, Int>>()
    val frequencyPositions = hashMapOf<Char, MutableSet<Pair<Int, Int>>>()
    for (i in matrix.indices) {
        for (j in matrix[i].indices) {
            val frequency = matrix[i][j]
            if (frequency == '.' || frequency == '#') {
                continue
            }
            if (frequency !in frequencyPositions) {
                frequencyPositions[frequency] = mutableSetOf(i to j)
            } else {
                frequencyPositions[frequency]!!.add(i to j)
            }
        }
    }
    for ((frequency, positions) in frequencyPositions) {
        antinodePositions.addAll(searchResonantAntinodes(positions.toList(), matrix.size, matrix[0].size))
    }
    return antinodePositions.count()
}

private fun searchAntinodes(positions: List<Pair<Int, Int>>, maxX: Int, maxY: Int): Set<Pair<Int, Int>> {
    val result = mutableSetOf<Pair<Int, Int>>()
    for (i in positions.indices) {
        for (j in positions.drop(i + 1).indices) {
            val dx = positions[i].first - positions[i + j + 1].first
            val dy = positions[i].second - positions[i + j + 1].second

            val x1 = positions[i].first + dx
            val x2 = positions[i + j + 1].first - dx

            val y1 = positions[i].second + dy
            val y2 = positions[i + j + 1].second - dy

            if (x1 in 0..<maxX && y1 in 0..<maxY) {
                result.add(x1 to y1)
            }
            if (x2 in 0..<maxX && y2 in 0..<maxY) {
                result.add(x2 to y2)
            }
        }
    }
    return result
}

private fun searchResonantAntinodes(positions: List<Pair<Int, Int>>, maxX: Int, maxY: Int): Set<Pair<Int, Int>> {
    val result = mutableSetOf<Pair<Int, Int>>()
    for (i in positions.indices) {
        for (j in positions.drop(i + 1).indices) {
            val dx = positions[i].first - positions[i + j + 1].first
            val dy = positions[i].second - positions[i + j + 1].second

            var x1 = positions[i].first + dx
            var x2 = positions[i + j + 1].first - dx

            var y1 = positions[i].second + dy
            var y2 = positions[i + j + 1].second - dy

            while (x1 in 0..<maxX && y1 in 0..<maxY) {
                result.add(x1 to y1)
                x1 += dx
                y1 += dy
            }

            while (x2 in 0..<maxX && y2 in 0..<maxY) {
                result.add(x2 to y2)
                x2 -= dx
                y2 -= dy
            }
        }
    }
    if (positions.size > 1) {
        result.addAll(positions)
    }
    return result
}