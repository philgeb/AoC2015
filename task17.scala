import scala.io.Source

object Main extends App {
  val total = 150
  val filename = "input.txt"

  val containers = Source.fromFile(filename).getLines.toList.map(_.toInt)
  val combs = for {
    len <- 1 to containers.length - 1
  } yield (0 to containers.length - 1).combinations(len).toList.map(_.map(i => containers(i)))

  val fittingCombs = combs.flatten.filter(_.sum == total)
  println("Part 1: " + fittingCombs.size)

  val minRequiredContainers = fittingCombs.min(Ordering.by((_:IndexedSeq[Int]).size)).size
  println("Part 2: " + fittingCombs.filter(_.size == minRequiredContainers).size)
}