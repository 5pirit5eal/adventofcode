import io.github.oshai.kotlinlogging.KotlinLogging
import java.io.File
import kotlin.math.log10
import kotlin.math.pow

private val logger = KotlinLogging.logger {}

fun seventhDay(): Pair<Long, Long> {
    val filePath = "inputs/seventh_day_input.txt"
    // read txt file and construct 2D matrix
    val lines = File(filePath).readLines()
    return getBinaryOperations(lines) to getTriOperations(lines)
}

class BinaryNode(val value: Long) {
    var left: BinaryNode? = null
    var right: BinaryNode? = null

    override fun toString(): String {
        return "$value"
    }
}

class TriNode(val value: Long) {
    var left: TriNode? = null
    var middle: TriNode? = null
    var right: TriNode? = null

    override fun toString(): String {
        return "$value"
    }
}


class BinaryTree(operands: List<Long>) {
    private val root: BinaryNode = BinaryNode(operands[0])

    init {
        for (operand in operands.drop(1)) {
            this.addNode(operand, this.root)
        }
    }

    private fun addNode(operand: Long, binaryNode: BinaryNode) {
        if (binaryNode.left == null) {
            binaryNode.left = BinaryNode(binaryNode.value * operand)
        } else {
            this.addNode(operand, binaryNode.left!!)
        }
        if (binaryNode.right == null) {
            binaryNode.right = BinaryNode(binaryNode.value + operand)
        } else {
            this.addNode(operand, binaryNode.right!!)
        }
    }

    fun findResult(equationResult: Long): Long {
        return this.findResultNode(equationResult, this.root)
    }

    private fun findResultNode(equationResult: Long, binaryNode: BinaryNode?): Long {
        if (binaryNode!!.left == null || binaryNode.right == null) {
            if (equationResult == binaryNode.value) {
                return equationResult
            } else {
                return 0
            }
        } else {
            logger.debug { "Searching for $equationResult, Looking at $binaryNode - left: ${binaryNode.left}, right: ${binaryNode.right}" }
            return when (equationResult) {
                this.findResultNode(equationResult, binaryNode.left) -> equationResult
                this.findResultNode(equationResult, binaryNode.right) -> equationResult
                else -> 0
            }
        }

    }
}

class TriTree(operands: List<Long>) {
    private val root: TriNode = TriNode(operands[0])

    init {
        for (operand in operands.drop(1)) {
            this.addNode(operand, this.root)
        }
    }

    private fun addNode(operand: Long, triNode: TriNode) {
        if (triNode.left == null) {
            triNode.left = TriNode(triNode.value * operand)
        } else {
            this.addNode(operand, triNode.left!!)
        }
        if (triNode.right == null) {
            triNode.right = TriNode(triNode.value + operand)
        } else {
            this.addNode(operand, triNode.right!!)
        }
        if (triNode.middle == null) {
            val magnitude = 10.0.pow(log10(operand.toDouble()).toInt() + 1).toLong()
            triNode.middle = TriNode(triNode.value * magnitude + operand)
        } else {
            this.addNode(operand, triNode.middle!!)
        }
    }

    fun findResult(equationResult: Long): Long {
        return this.findResultNode(equationResult, this.root)
    }

    private fun findResultNode(equationResult: Long, triNode: TriNode?): Long {
        if (triNode!!.left == null || triNode.right == null || triNode.middle == null) {
            if (equationResult == triNode.value) {
                return equationResult
            } else {
                return 0
            }
        } else {
            logger.debug { "Searching for $equationResult, Looking at $triNode - left: ${triNode.left}, middle: ${triNode.middle}, right: ${triNode.right}" }
            return when (equationResult) {
                this.findResultNode(equationResult, triNode.left) -> equationResult
                this.findResultNode(equationResult, triNode.middle) -> equationResult
                this.findResultNode(equationResult, triNode.right) -> equationResult
                else -> 0
            }
        }

    }
}

private fun getBinaryOperations(equations: List<String>): Long {
    // split by :
    var result: Long = 0
    for (equation in equations) {
        val (rawEquationResult, rawOperands) = equation.split(":")
        val equationResult = rawEquationResult.toLong()
        val operands = rawOperands.split(" ").drop(1).map { it -> it.toLong() }
        val tree = BinaryTree(operands)
        val searchResult = tree.findResult(equationResult)
        logger.info { "Found $searchResult, expected $equationResult" }
        result += searchResult

    }
    return result
}

private fun getTriOperations(equations: List<String>): Long {
    // split by :
    var result: Long = 0
    for (equation in equations) {
        val (rawEquationResult, rawOperands) = equation.split(":")
        val equationResult = rawEquationResult.toLong()
        val operands = rawOperands.split(" ").drop(1).map { it -> it.toLong() }
        val tree = TriTree(operands)
        val searchResult = tree.findResult(equationResult)
        logger.info { "Found $searchResult, expected $equationResult" }
        result += searchResult

    }
    return result
}

