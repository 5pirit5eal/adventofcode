import io.github.oshai.kotlinlogging.KotlinLogging
import java.io.File
import java.rmi.UnexpectedException
import java.security.InvalidParameterException

private val logger = KotlinLogging.logger {}

fun sixthDay(): Int {
    val filePath = "inputs/sixth_day_input.txt"
    // read txt file and construct 2D matrix
    val lines = File(filePath).readLines()
    val matrix = Array(lines.size) { CharArray(lines[0].length) }

    for (i in lines.indices) {
        for (j in lines[i].indices) {
            matrix[i][j] = lines[i][j]
        }
    }
    return countObstructions(matrix)
}

enum class Direction(val orientation: Char) {
    UPWARDS('^'), RIGHT('>'), DOWNWARDS('v'), LEFT('<');

    fun next(): Direction {
        val values = Direction.entries
        val nextIndex = (this.ordinal + 1) % values.size
        return values[nextIndex]
    }

    companion object {
        fun fromChar(orientation: Char): Direction? {
            return entries.find { it.orientation == orientation }
        }
    }
}


class Guard(var position: Pair<Int, Int>, var direction: Direction) {
    var inside = true
    var tries = 0

    init {
        if (position.first < 0 || position.second < 0) {
            throw InvalidParameterException("The position $position contains negative values.")
        }
    }

    /**
     * Does the next step for a given matrix. Updates position property
     *
     * @param matrix the matrix to step through
     * @return true, if the guard stays within the matrix, false otherwise
     */
    fun nextStep(matrix: Array<CharArray>): Boolean {
        val nextPosition = this.getNextPosition()
        if (isOutside(matrix, nextPosition)) {
            logger.debug { "Outside of matrix at $position, wanting to go to $nextPosition for matrix of [${matrix.size},${matrix[0].size}]" }
            return false
        } else if (tries > 5) {
            logger.error { "I am trapped at $position!" }
            return false
        } else if (isObstacle(matrix, nextPosition)) {
            logger.debug { "Found an obstacle at $nextPosition, retries at $tries" }
            direction = direction.next()
            tries++
            return this.nextStep(matrix)
        } else {
            position = nextPosition
            tries = 0
            return true
        }
    }

    private fun isObstacle(matrix: Array<CharArray>, nextPosition: Pair<Int, Int>): Boolean {
        return (matrix[nextPosition.first][nextPosition.second] in (listOf(
            '#', 'O'
        ) + Direction.entries.map { it.orientation }))
    }

    private fun isOutside(matrix: Array<CharArray>, position: Pair<Int, Int>): Boolean {
        return (position.first < 0 || position.second < 0 || position.first > matrix.size - 1 || position.second > matrix[position.first].size - 1)
    }

    private fun getNextPosition(): Pair<Int, Int> {
        return when (this.direction) {
            Direction.UPWARDS -> Pair(this.position.first - 1, this.position.second)
            Direction.DOWNWARDS -> Pair(this.position.first + 1, this.position.second)
            Direction.RIGHT -> Pair(this.position.first, this.position.second + 1)
            Direction.LEFT -> Pair(this.position.first, this.position.second - 1)
        }
    }
}

fun getUniquePositions(matrix: Array<CharArray>): Set<Pair<Int, Int>> {
    val uniquePositions = mutableSetOf<Pair<Int, Int>>()
    var guard: Guard? = null
    outer@ for (i in matrix.indices) {
        for (j in matrix[i].indices) {
            if (matrix[i][j] in Direction.entries.map { it.orientation }) {
                guard = Guard(Pair(i, j), Direction.fromChar(matrix[i][j])!!)
                uniquePositions.add(Pair(i, j))
                break@outer
            }
        }
    }
    if (guard == null) {
        throw UnexpectedException("Guard should be defined in matrix!")
    }
    logger.info { "Starting guard patrols..." }
    while (guard.inside) {
        if (!guard.nextStep(matrix)) {
            guard.inside = false
        }
        uniquePositions.add(guard.position)
    }
    return uniquePositions.toSet()
}

fun countObstructions(matrix: Array<CharArray>): Int {
    var firstGuard: Guard? = null
    var startingPosition: Triple<Int, Int, Direction>? = null
    outer@ for (i in matrix.indices) {
        for (j in matrix[i].indices) {
            if (matrix[i][j] in Direction.entries.map { it.orientation }) {
                firstGuard = Guard(Pair(i, j), Direction.fromChar(matrix[i][j])!!)
                startingPosition = Triple(i, j, Direction.fromChar(matrix[i][j])!!)
                break@outer
            }
        }
    }
    if (firstGuard == null || startingPosition == null) {
        throw UnexpectedException("Guard should be defined in matrix!")
    }
    logger.info { "Starting testing obstructions..." }
    var obstructionCount = 0
    val obstructionOptions = getUniquePositions(matrix)
    for (obstruction in obstructionOptions) {
        val previousObject = matrix[obstruction.first][obstruction.second]
        if (previousObject in listOf('#', '^')) {
            continue
        }
        matrix[obstruction.first][obstruction.second] = 'O'
        // init guard
        val guard = Guard(firstGuard.position, firstGuard.direction)
        val uniquePositions = mutableSetOf(startingPosition)
        var obstructionFound = false

        logger.debug { "Starting new run with obstruction at $obstruction" }
        while (guard.inside && !obstructionFound) {
            val position = guard.position
            val isInside = guard.nextStep(matrix)
            if (isInside) {
                val newlyWalked =
                    uniquePositions.add(Triple(guard.position.first, guard.position.second, guard.direction))
                if (!newlyWalked) {
                    logger.debug { "Obstruction found at ${guard.position}" }
                    obstructionFound = true
                    obstructionCount++
                }
            } else {
                guard.inside = false
            }
            matrix[position.first][position.second] = '.'

        }
        // reset matrix
        matrix[obstruction.first][obstruction.second] = previousObject

    }



    return obstructionCount
}
