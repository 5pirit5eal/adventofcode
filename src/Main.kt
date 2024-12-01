import measureDistance
import measureSimilarity
import kotlin.system.measureTimeMillis

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
fun main() {
    val (a, b) = parseInput()
    val duration = measureTimeMillis {
        println(measureDistance(a, b))
        println(measureSimilarity(a, b))
    }
    println("Time taken $duration ms")

}


