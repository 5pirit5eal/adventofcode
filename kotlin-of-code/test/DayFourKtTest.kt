import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.MethodSource
import java.io.File

class DayFourKtTest {
    companion object {
        @JvmStatic
        fun provideTestCases(): List<Any> {
            var filePath = "test/fourthtest.txt"
            var lines = File(filePath).readLines()
            val matrixLarge = Array(lines.size) { CharArray(lines[0].length) }

            for (i in lines.indices) {
                for (j in lines[i].indices) {
                    matrixLarge[i][j] = lines[i][j]
                }
            }
            filePath = "test/fourthtestSmaller.txt"
            lines = File(filePath).readLines()
            val matrixSmall = Array(lines.size) { CharArray(lines[0].length) }

            for (i in lines.indices) {
                for (j in lines[i].indices) {
                    matrixSmall[i][j] = lines[i][j]
                }
            }
            val fileString = "XMASAMX"
            val matrixSmaller = Array(1) { CharArray(fileString.length) }

            for (j in fileString.indices) {
                matrixSmaller[0][j] = fileString[j]
            }


            return listOf(arrayOf(matrixSmall, 4), arrayOf(matrixLarge, 18), arrayOf(matrixSmaller, 2))
        }
    }

    @ParameterizedTest
    @MethodSource("provideTestCases")
    fun testCountXMAS(input: Array<CharArray>, expected: Int) {
        val result = countXMAS(input)
        assertEquals(expected, result)
    }
}

class DayFourKtTest2 {
    companion object {
        @JvmStatic
        fun provideTestCases(): List<Any> {
            var filePath = "test/fourthtest.txt"
            var lines = File(filePath).readLines()
            val matrixLarge = Array(lines.size) { CharArray(lines[0].length) }

            for (i in lines.indices) {
                for (j in lines[i].indices) {
                    matrixLarge[i][j] = lines[i][j]
                }
            }
            filePath = "test/fourthtestSmaller.txt"
            lines = File(filePath).readLines()
            val matrixSmall = Array(lines.size) { CharArray(lines[0].length) }

            for (i in lines.indices) {
                for (j in lines[i].indices) {
                    matrixSmall[i][j] = lines[i][j]
                }
            }
            val fileString = "M.S\n" +
                    ".A.\n" +
                    "M.S"
            val matrixSmaller = Array(3) { CharArray(3) }
            for ((i, split) in fileString.split("\n").withIndex()) {
                for (j in split.indices) {
                    matrixSmaller[i][j] = split[j]
                }
            }



            return listOf(arrayOf(matrixSmall, 0), arrayOf(matrixLarge, 9), arrayOf(matrixSmaller, 1))
        }
    }

    @ParameterizedTest
    @MethodSource("provideTestCases")
    fun testCountMAS(input: Array<CharArray>, expected: Int) {
        val result = countMAS(input)
        assertEquals(expected, result)
    }
}