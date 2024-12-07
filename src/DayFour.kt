import io.github.oshai.kotlinlogging.KotlinLogging
import java.io.File
import java.rmi.UnexpectedException

private val logger = KotlinLogging.logger {}

fun fourthDay(): Pair<Int, Int> {
    val filePath = "inputs/fourth_day_input.txt"
    // read txt file and construct 2D matrix
    val lines = File(filePath).readLines()
    val matrix = Array(lines.size) { CharArray(lines[0].length) }

    for (i in lines.indices) {
        for (j in lines[i].indices) {
            matrix[i][j] = lines[i][j]
        }
    }

    return Pair(countXMAS(matrix), countMAS(matrix))
}

fun getDiagonal(matrix: Array<CharArray>, startRow: Int, startColumn: Int): String {
    val rows = matrix.size
    val cols = matrix[0].size
    var diagonal = ""
    var r = startRow
    var c = startColumn

    while (r < rows && c < cols) {
        diagonal += matrix[r][c]
        r++
        c++
    }
    return diagonal
}

fun getReverseDiagonal(matrix: Array<CharArray>, startRow: Int, startColumn: Int): String {
    var diagonal = ""
    var r = startRow
    var c = startColumn

    while (r >= 0 && c < matrix[0].size) {
        diagonal += matrix[r][c]
        r--
        c++
    }
    return diagonal
}

fun countXMAS(matrix: Array<CharArray>): Int {
    // Check all directions for XMAS and SAMX
    var result = 0
    for (pattern in arrayOf("XMAS".toRegex(), "SAMX".toRegex())) {
        for (i in matrix.indices) {
            result += pattern.findAll(matrix[i].joinToString("")).count()
            logger.debug { "For row $i, we increased to $result on horizontal" }
            logger.debug { "This is the line ${matrix[i].joinToString("")}" }
            result += pattern.findAll(getDiagonal(matrix = matrix, startColumn = 0, startRow = i)).count()
            logger.debug { "For row $i, we increased to $result on diagonal" }
            logger.debug { "This is the line ${getDiagonal(matrix = matrix, startColumn = 0, startRow = i)}" }
            result += pattern.findAll(
                getReverseDiagonal(
                    matrix = matrix,
                    startColumn = 0,
                    startRow = i
                )
            ).count()
            logger.debug { "For row $i, we increased to $result on reverse diagonal" }
            logger.debug {
                "This is the line ${
                    getReverseDiagonal(
                        matrix = matrix,
                        startColumn = 0,
                        startRow = i
                    )
                }"
            }

        }
        for (j in matrix[0].indices) {
            result += pattern.findAll(matrix.map { it[j] }.joinToString("")).count()
            logger.debug { "For column $j, we increased to $result on vertical" }
            logger.debug { "This is the line ${matrix.map { it[j] }.joinToString("")}" }
            // Skip as to not look at the diagonal at 0,0 twice
            if (j == 0) {
                continue
            }
            result += pattern.findAll(getDiagonal(matrix = matrix, startColumn = j, startRow = 0)).count()
            logger.debug { "For column $j, we increased to $result on diagonal" }
            logger.debug { "This is the line ${getDiagonal(matrix = matrix, startColumn = j, startRow = 0)}" }
            result += pattern.findAll(getReverseDiagonal(matrix = matrix, startColumn = j, startRow = matrix.size - 1))
                .count()
            logger.debug { "For column $j, we increased to $result on reverse diagonal" }
            logger.debug {
                "This is the line ${
                    getReverseDiagonal(
                        matrix = matrix, startColumn = j, startRow = matrix.size - 1
                    )
                }"
            }
        }
    }
    return result
}

fun countMAS(matrix: Array<CharArray>): Int {
    // Check all diagonals for MAS and SAM and create a matrix where all As are counted
    // When As are 2 then increase the result counter
    val AMatrix = Array(matrix.size) { IntArray(matrix[0].size) }
    var result = 0
    for (pattern in arrayOf("MAS".toRegex(), "SAM".toRegex())) {
        for (i in matrix.indices) {
            val diagonalHits = pattern.findAll(getDiagonal(matrix = matrix, startColumn = 0, startRow = i))
            logger.debug { "This is the diagonal line ${getDiagonal(matrix = matrix, startColumn = 0, startRow = i)}" }
            for (hit in diagonalHits) {
                val pos = (hit.range.first + 1)
                AMatrix[i + pos][pos]++
                logger.debug { "For row $i, we increased by ${AMatrix[i + pos][pos]} on diagonal" }
            }


            val reverseHits = pattern.findAll(
                getReverseDiagonal(
                    matrix = matrix,
                    startColumn = 0,
                    startRow = i
                )
            )
            logger.debug {
                "This is the reverse diagonal line ${
                    getReverseDiagonal(
                        matrix = matrix,
                        startColumn = 0,
                        startRow = i
                    )
                }"
            }
            for (hit in reverseHits) {
                val pos = (hit.range.first + 1)
                AMatrix[i - pos][pos]++
                logger.debug { "For row $i, we increased by ${AMatrix[i - pos][pos]} on reverse diagonal" }
            }


        }
        for (j in 1..<matrix[0].size) {
            val diagonalHitsColumn = pattern.findAll(getDiagonal(matrix = matrix, startColumn = j, startRow = 0))
            logger.debug { "This is the diagonal line ${getDiagonal(matrix = matrix, startColumn = j, startRow = 0)}" }
            for (hit in diagonalHitsColumn) {
                val pos = (hit.range.first + 1)
                AMatrix[pos][j + pos]++
                logger.debug { "For column $j, we increased by ${AMatrix[pos][j + pos]} on reverse diagonal" }
            }

            val reverseHitsColumn =
                pattern.findAll(getReverseDiagonal(matrix = matrix, startColumn = j, startRow = matrix.size - 1))
            logger.debug {
                "This is the reverse diagonal line ${
                    getReverseDiagonal(
                        matrix = matrix, startColumn = j, startRow = matrix.size - 1
                    )
                }"
            }
            for (hit in reverseHitsColumn) {
                val pos = (hit.range.first + 1)
                AMatrix[matrix.size - 1 - pos][j + pos]++
                logger.debug { "For column $j, we increased by ${AMatrix[matrix.size - 1 - pos][j + pos]} on reverse diagonal" }
            }

        }
    }
    for (row in AMatrix) {
        for (col in row) {
            logger.debug { col }
        }
        logger.debug {}
    }
    for (i in AMatrix.indices) {
        for (j in AMatrix[i].indices) {
            when (AMatrix[i][j]) {
                3, 4, 5, 6 -> {
                    logger.debug { "This should not happen" }
                    throw UnexpectedException("The value 3 was unexpected at position $i, $j")
                }

                2 -> {
                    result++
                    logger.debug { "For $i,$j we increased to $result" }
                }
            }
        }
    }
    return result
}

