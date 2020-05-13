import scala.io.Source
import scala.collection.mutable.HashMap

object Main extends App {
  val input = 34000000

  // Part 1
  val converted_in = input / 10

  var houses = Array.fill(converted_in)(0)

  for (
    i <- 1 until converted_in;
    j <- i until converted_in by i
  ) {
    houses(j) += i
  }

  println("Part 1: " + houses.zipWithIndex.filter(_._1 > converted_in).minBy(_._2)._2)

  // Part 2
  houses = Array.fill(converted_in)(0)

  for (
    i <- 1 until converted_in;
    j <- i until converted_in by i if j < i * 50
  ) {
    houses(j) += i * 11
  }

  println("Part 2: " + houses.zipWithIndex.filter(_._1 > input).minBy(_._2)._2)
}