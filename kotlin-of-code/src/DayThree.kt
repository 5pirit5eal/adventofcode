import java.io.File

fun thirdDay(): Int {
    val filePath = "inputs/third_day_input.txt"
    var result = 0
    // read txt file
    // iterate and append to output lists
    File(filePath).useLines { lines ->
        result += splitDosDonts(lines.joinToString(""))
    }
    return result
}

fun splitDosDonts(corruptData: String): Int {
    val dontParts = corruptData.split("don't()")
    var result = regexMul(dontParts[0])

    for (dontPart in dontParts.drop(1)) {
        val doParts = dontPart.split("do()")
        result += regexMul(doParts.drop(1).joinToString(separator = ","))
    }
    return result
}


fun regexMul(corruptData: String): Int {
    var result = 0
    val mulPattern = "mul\\((\\d+),(\\d+)\\)".toRegex()
    for (foundMul in mulPattern.findAll(corruptData)) {
        result += foundMul.groupValues[1].toInt() * foundMul.groupValues[2].toInt()
    }
    return result
}