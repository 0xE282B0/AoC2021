class FishCluster(internal val num:Long, age:Int=8) {
    private var fish: Fish

    init {
        this.fish = Fish(age)
    }

    fun nextDay() : FishCluster? {
        if (fish.nextDay())
            return FishCluster(num)
        else
            return null
    }

    override fun toString(): String {
        return "$num $fish"
    }
}

fun main() {
    val input = {}::class.java.getResource("data/input.txt").readText().trim().split(",")

    val cluster = arrayListOf<FishCluster>()
    input.groupBy { it }.forEach{ cluster.add(FishCluster(it.value.size+0L,Integer.parseInt(it.key))) }

    for ( i in 1..256){
        val newFish = arrayListOf<FishCluster?>()
        for(c in cluster){
            newFish.add(c.nextDay())
        }
        var sum = newFish.filterNotNull().map { it.num }.sum()
        if(sum > 0)
            cluster.add(FishCluster(sum))
    }
    println(cluster.map{it.num+0L}.sum())
}