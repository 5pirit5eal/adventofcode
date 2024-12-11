import io.github.oshai.kotlinlogging.KotlinLogging
import java.io.File

private val logger = KotlinLogging.logger {}

fun seventhDay(): Long {
    val filePath = "inputs/seventh_day_input.txt"
    // read txt file and construct 2D matrix
    val lines = File(filePath).readLines()

    return getOperations(lines)
}

class Node(val value: Long) {
    var left: Node? = null
    var right: Node? = null

    override fun toString(): String {
        return "$value"
    }
}

class BinaryTree(operands: List<Long>) {
    private val root: Node = Node(operands[0])

    init {
        for (operand in operands.drop(1)) {
            this.addNode(operand, this.root)
        }
    }

    private fun addNode(operand: Long, node: Node) {
        if (node.left == null) {
            node.left = Node(node.value * operand)
        } else {
            this.addNode(operand, node.left!!)
        }
        if (node.right == null) {
            node.right = Node(node.value + operand)
        } else {
            this.addNode(operand, node.right!!)
        }
    }

    fun findResult(equationResult: Long): Long {
        return this.findResultNode(equationResult, this.root)
    }

    private fun findResultNode(equationResult: Long, node: Node?): Long {
        if (node!!.left == null || node.right == null) {
            if (equationResult == node.value) {
                return equationResult
            } else {
                return 0
            }
        } else {
            logger.debug { "Searching for $equationResult, Looking at $node - left: ${node.left}, right: ${node.right}" }
            if (equationResult == this.findResultNode(equationResult, node.left)) {
                return equationResult
            } else if (equationResult == this.findResultNode(equationResult, node.right)) {
                return equationResult
            }
            return 0
        }

    }
}

private fun getOperations(equations: List<String>): Long {
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

