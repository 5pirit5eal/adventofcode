import java.time.LocalDate
import kotlin.system.measureTimeMillis

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
fun main() {
    val duration = measureTimeMillis {
        advent()
    }
    println("Time taken $duration ms")

}

fun advent() {
    when (LocalDate.now().dayOfMonth) {
        1 -> {
            val (a, b) = parseInput()
            println(measureDistance(a, b))
            println(measureSimilarity(a, b))
        }

        2 -> println(secondDay())
        3 -> println(thirdDay())
        4 -> println(fourthDay())
        5 -> println(fifthDay())
        6 -> println(sixthDay())
        7 -> println(seventhDay())
        8 -> println(eightDay())
        9 -> println(ninthDay())
        17 -> println(tenthDay())
    }
}
