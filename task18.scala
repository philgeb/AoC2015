import scala.io.Source

object Main extends App {
  var size = 100
  val steps = 100
  val filename = "input.txt"

  var init_rows = 
    Array(("." * (size + 2)).toArray) ++ 
    Source.fromFile(filename).getLines.toArray.map(xs => ('.' + xs + '.').toArray) ++ 
    Array(("." * (size + 2)).toArray)

  // part 1
  var rows = init_rows.map(_.clone())
  var next_rows = rows.map(_.clone())

  for (_ <- 0 until steps) {
    for (i <- 1 to size; j <- 1 to size) {
      val sumOn = Array(
        rows(i-1)(j-1), rows(i-1)(j), rows(i-1)(j+1),
        rows(i)(j-1), rows(i)(j+1),
        rows(i+1)(j-1), rows(i+1)(j), rows(i+1)(j+1)
      ).filter(_ == '#').size

      if (rows(i)(j) == '#' && sumOn != 2 && sumOn != 3) {
        next_rows(i)(j) = '.'
      } else if (rows(i)(j) == '.' && sumOn == 3) {
        next_rows(i)(j) = '#'
      }
    }

    rows = next_rows.map(_.clone())
  }
  
  println("Part 1: " + rows.flatten.filter(_ == '#').size)

  // part 2
  rows = init_rows.map(_.clone())
  rows(1)(1) = '#'
  rows(1)(size) = '#'
  rows(size)(1) = '#'
  rows(size)(size) = '#'
  next_rows = rows.map(_.clone())

  val corners = Array((1,1), (1,size), (size,1), (size,size))

  for (_ <- 0 until steps) {
    for (i <- 1 to size; j <- 1 to size if !(corners contains (i,j))) {
      val sumOn = Array(
        rows(i-1)(j-1), rows(i-1)(j), rows(i-1)(j+1),
        rows(i)(j-1), rows(i)(j+1),
        rows(i+1)(j-1), rows(i+1)(j), rows(i+1)(j+1)
      ).filter(_ == '#').size

      if (rows(i)(j) == '#' && sumOn != 2 && sumOn != 3) {
        next_rows(i)(j) = '.'
      } else if (rows(i)(j) == '.' && sumOn == 3) {
        next_rows(i)(j) = '#'
      }
    }

    rows = next_rows.map(_.clone())
  }

  println("Part 2: " + rows.flatten.filter(_ == '#').size)
}