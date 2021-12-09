class Fish() {
    private var cycle: Int = 8

    constructor(age: Int) : this() {
        cycle = age
    }

    fun nextDay(): Boolean{
        cycle--
        if (cycle < 0){
            cycle = 6
            return true
        }
        return false
    }

    override fun toString(): String {
        return "Fish: $cycle"
    }
}

fun main() {
    val input = {}::class.java.getResource("data/input.txt").readText().trim().split(",")

    val fish = arrayListOf<Fish>()
    for(i in input){
        fish.add(Fish(Integer.parseInt(i)))
    }
//    println(fish)

    for ( i in 1..80){
        val newFish = arrayListOf<Fish>()
        for(f in fish){
            if(f.nextDay())
                newFish.add(Fish())
        }
        fish.addAll(newFish)
    }
    print("Total number: ${fish.size}")
}