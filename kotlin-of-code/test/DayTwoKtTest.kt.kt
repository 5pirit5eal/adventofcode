import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.MethodSource

class DayTwoKtTest {
    companion object {
        @JvmStatic
        fun provideTestCases() = listOf(
            arrayOf(listOf(48, 46, 47, 49, 51, 54, 56), true),
            arrayOf(mutableListOf(1, 1, 2, 3, 4, 5), true),
            arrayOf(mutableListOf(1, 2, 3, 4, 5, 5), true),
            arrayOf(mutableListOf(5, 1, 2, 3, 4, 5), true),
            arrayOf(mutableListOf(1, 4, 3, 2, 1), true),
            arrayOf(mutableListOf(1, 6, 7, 8, 9), true),
            arrayOf(mutableListOf(1, 2, 3, 4, 3), true),
            arrayOf(mutableListOf(9, 8, 7, 6, 7), true),
            arrayOf(mutableListOf(7, 10, 8, 10, 11), true),
            arrayOf(mutableListOf(7, 10, 8, 10, 11), true),
            arrayOf(mutableListOf(29, 28, 27, 25, 26, 25, 22, 20), true),
        )
    }

    @ParameterizedTest
    @MethodSource("provideTestCases")
    fun testProblemDampener(input: MutableList<Int>, expected: Boolean) {
        val result = problemDampener(input)
        assertEquals(expected, result)
    }
}