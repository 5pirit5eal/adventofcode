import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.MethodSource

class DayThreeHelperKtTest {
    companion object {
        @JvmStatic
        fun provideTestCases() = listOf(
            arrayOf("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161),
            arrayOf("mul( 1, 1)", 0),
            arrayOf("mul(1*,1)", 0),
            arrayOf("_mul(1,5)", 5),
            arrayOf("mul(-1,5)", 0),
            arrayOf("mul(1,-5)", 0),
            arrayOf("_mul(1, 1)L", 0),
        )
    }

    @ParameterizedTest
    @MethodSource("provideTestCases")
    fun testRegexMul(input: String, expected: Int) {
        val result = regexMul(input)
        assertEquals(expected, result)
    }
}

class DayThreeKtTest {
    companion object {
        @JvmStatic
        fun provideTestCases() = listOf(
            arrayOf(
                "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48
            ),
            arrayOf(
                "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))don't()_mul(5,5)+mul(32,64](mul(11,8)",
                96
            ),
        )
    }

    @ParameterizedTest
    @MethodSource("provideTestCases")
    fun testSplitDosDonts(input: String, expected: Int) {
        val result = splitDosDonts(input)
        assertEquals(expected, result)
    }
}