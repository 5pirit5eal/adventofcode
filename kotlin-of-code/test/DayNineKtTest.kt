import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.MethodSource

class DayNineKtTest {
    companion object {
        @JvmStatic
        fun provideTestCases(): List<Any> {
            return listOf(
                arrayOf("2333133121414131402", "00...111...2...333.44.5555.6666.777.888899")
            )
        }
    }

    @ParameterizedTest
    @MethodSource("provideTestCases")
    fun testSparsify(input: String, expected: String) {
        val result = sparsify(input)
        assertEquals(expected, result)
    }

    @Test
    fun testRearrange() {
        val input = "00...111...2...333.44.5555.6666.777.888899"
        val expected = "0099811188827773336446555566.............."
        val result = rearrange(input.map { it.toString() }.toMutableList())
        assertEquals(expected, result)
    }

    @Test
    fun testBlockRearrange() {
        val input = "2333133121414131402"
        val expected: Long = 2858
        assertEquals(expected, blockRearrange(input))
    }
}

