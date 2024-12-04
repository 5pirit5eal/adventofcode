import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.MethodSource
import java.io.File

fun xmasDetection(matrix: Array<CharArray>): Int {
    // Check all directions for XMAS and SAMX
    var result = 0
    val pattern = "(XMAS|SAMX)".toRegex()
    for (i in matrix.indices) {
        result += pattern.findAll(matrix[i].joinToString("")).count()

        result += pattern.findAll(getDiagonal(matrix = matrix, startColumn = 0, startRow = i)).count()
        result += pattern.findAll(getReverseDiagonal(matrix = matrix, startColumn = matrix[0].size - 1, startRow = i))
            .forEachIndexed({index, matchResult -> matchResult.groupValues .forEach[]})
    }
    for (j in matrix[0].indices) {
        result += pattern.findAll(matrix.map { it[j] }.joinToString("")).count()
        result += pattern.findAll(getDiagonal(matrix = matrix, startColumn = j, startRow = 0)).count()
        result += pattern.findAll(getReverseDiagonal(matrix = matrix, startColumn = j, startRow = matrix.size - 1))
            .count()
    }
    return result
}

class `DayFourKtTest.kt` {
    companion object {
        @JvmStatic
        fun provideTestCases(): List<Any> {
            val filePath = "test/fourthtest.txt"
            val lines = File(filePath).readLines()
            val matrix = Array(lines.size) { CharArray(lines[0].length) }

            for (i in lines.indices) {
                for (j in lines[i].indices) {
                    matrix[i][j] = lines[i][j]
                }
            }

            return listOf(arrayOf(matrix, "."))
        }
    }

    @ParameterizedTest
    @MethodSource("provideTestCases")
    fun testXMASDetection(input: Array<CharArray>, expected: String) {
        val matrix = xmasDetection()
        for (i in matrix.indices) {
            for (j in matrix[i].indices) {
                assertEquals(matrix[i][j], expected)
            }
        }
        assertEquals(expected, result)
    }
}