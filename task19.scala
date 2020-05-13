import scala.io.Source
import scala.collection.mutable.HashMap

object Main extends App {
  val filename = "input.txt"

  var initial = new String
  var replacments = List[(String, String)]()
  val reg = "(.*) => (.*)".r
  for (line <- Source.fromFile(filename).getLines) {
    line match {
      case reg(lhs, rhs) => replacments :+= ((lhs, rhs))
      case _ => if (!line.isEmpty()) initial = line
    }
  }

  // Part 1
  val results = collection.mutable.Set[String]()
  for (
    (repl_l, repl_r) <- replacments;
    i <- 0 until initial.size) 
  {
    val (left, right) = initial.splitAt(i)
    val res = left + right.replaceFirst(repl_l, repl_r)
    results += res
  }
  // above algorithm might also create a result that didn't replac anything, so remove this case here
  results -= initial

  println("Part 1: " + results.size)
}