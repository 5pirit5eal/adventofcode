import java.io.File

fun fourthDay(): Int {
    val filePath = "inputs/fourth_day_input.txt"
    var result = 0
    // read txt file
    // iterate and append to output lists
    val lines = File(filePath).readLines()
    val matrix = Array(lines.size) { CharArray(lines[0].length) }

    for (i in lines.indices) {
        for (j in lines[i].indices) {
            matrix[i][j] = lines[i][j]
        }
    }

    return countXMAS(matrix)
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

    while (r >= 0 && c >= 0) {
        diagonal += matrix[r][c]
        r--
        c--
    }
    return diagonal
}

fun countXMAS(matrix: Array<CharArray>): Int {
    // Check all directions for XMAS and SAMX
    var result = 0
    val pattern = "(XMAS|SAMX)".toRegex()
    for (i in matrix.indices) {
        result += pattern.findAll(matrix[i].joinToString("")).count()
        
        result += pattern.findAll(getDiagonal(matrix = matrix, startColumn = 0, startRow = i)).count()
        result += pattern.findAll(getReverseDiagonal(matrix = matrix, startColumn = matrix[0].size - 1, startRow = i))
            .count()
    }
    for (j in matrix[0].indices) {
        result += pattern.findAll(matrix.map { it[j] }.joinToString("")).count()
        result += pattern.findAll(getDiagonal(matrix = matrix, startColumn = j, startRow = 0)).count()
        result += pattern.findAll(getReverseDiagonal(matrix = matrix, startColumn = j, startRow = matrix.size - 1))
            .count()
    }
    return result
}