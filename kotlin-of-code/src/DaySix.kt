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
    // return getUniquePositions(matrix).distinctBy { Pair(it.first, it.second) }.count()
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
    private var tries = 0

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
    fun nextStep(matrix: Array<CharArray>) {
        val nextPosition = this.getNextPosition()
        if (isOutside(matrix, nextPosition)) {
            logger.debug { "Outside of matrix at $position, wanting to go to $nextPosition for matrix of [${matrix.size},${matrix[0].size}]" }
            inside = false
        } else if (tries > 4) {
            logger.error { "I am trapped at $position!" }
            inside = false
        } else if (isObstacle(matrix, nextPosition)) {
            logger.debug { "Found an obstacle at $nextPosition, retries at $tries" }
            direction = direction.next()
            logger.debug { "Moving to $direction" }
            tries++
            this.nextStep(matrix)
        } else {
            position = nextPosition
            tries = 0
        }
    }

    private fun isObstacle(matrix: Array<CharArray>, nextPosition: Pair<Int, Int>): Boolean {
        return (matrix[nextPosition.first][nextPosition.second] in listOf('#', 'O'))
    }

    private fun isOutside(matrix: Array<CharArray>, position: Pair<Int, Int>): Boolean {
        return (position.first < 0 || position.second < 0 || position.first > matrix.size - 1 || position.second > matrix[position.first].size - 1)
    }

    private fun getNextPosition(): Pair<Int, Int> {
        return when (direction) {
            Direction.UPWARDS -> Pair(position.first - 1, position.second)
            Direction.DOWNWARDS -> Pair(position.first + 1, position.second)
            Direction.RIGHT -> Pair(position.first, position.second + 1)
            Direction.LEFT -> Pair(position.first, position.second - 1)
        }
    }
}

fun getUniquePositions(matrix: Array<CharArray>): List<Triple<Int, Int, Direction>> {
    val uniquePositions = mutableSetOf<Triple<Int, Int, Direction>>()
    var guard: Guard? = null
    outer@ for (i in matrix.indices) {
        for (j in matrix[i].indices) {
            if (matrix[i][j] in Direction.entries.map { it.orientation }) {
                guard = Guard(Pair(i, j), Direction.fromChar(matrix[i][j])!!)
                uniquePositions.add(Triple(i, j, guard.direction))
                break@outer
            }
        }
    }
    if (guard == null) {
        throw UnexpectedException("Guard should be defined in matrix!")
    }
    logger.info { "Starting guard patrols..." }
    while (guard.inside) {
        guard.nextStep(matrix)
        logger.debug { "Adding ${Triple(guard.position.first, guard.position.second, guard.direction)}" }
        uniquePositions.add(Triple(guard.position.first, guard.position.second, guard.direction))
    }
    return uniquePositions.toList()
}

fun countObstructions(matrix: Array<CharArray>): Int {
    logger.info { "Starting testing obstructions..." }
    val obstructionPositions = mutableSetOf<Pair<Int, Int>>()
    val obstructionOptions = getUniquePositions(matrix).distinctBy { it.first to it.second }
    for ((o, obstructionPosition) in obstructionOptions.withIndex()) {
        if (o < 1) {
            continue
        }
        val previousObject = matrix[obstructionPosition.first][obstructionPosition.second]
        matrix[obstructionPosition.first][obstructionPosition.second] = 'O'
        // reinit guard
        val guard =
            Guard(obstructionOptions[o - 1].first to obstructionOptions[o - 1].second, obstructionOptions[o - 1].third)
        val uniquePositions = obstructionOptions.subList(0, o + 1).toMutableSet()
        var obstructionFound = false

        logger.debug { "Starting new run with guard at ${guard.position} and obstruction at $obstructionPosition" }
        while (guard.inside && !obstructionFound) {
            guard.nextStep(matrix)
            if (guard.inside) {
                val newlyWalked =
                    uniquePositions.add(Triple(guard.position.first, guard.position.second, guard.direction))
                if (!newlyWalked) {
                    logger.debug { "Repetition found at ${guard.position} for obstruction at $obstructionPosition" }
                    obstructionFound = true
                    obstructionPositions.add(Pair(obstructionPosition.first, obstructionPosition.second))
                }
            }
        }
        // reset matrix
        matrix[obstructionPosition.first][obstructionPosition.second] = previousObject
    }
    return obstructionPositions.count()
}
