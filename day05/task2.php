<?php
$handle = fopen("data/input.txt", "r");

function isHorizontal($x1,$y1,$x2,$y2){
    return $x1==$x2;
}

function isVertical($x1,$y1,$x2,$y2){
    return $y1==$y2;
}

function addPoint(&$grid,$x,$y){
    #print("x".$x."y".$y." \n");
    if (!array_key_exists($x,$grid)){
        $grid[$x] = [];
    }
    if (!array_key_exists($y,$grid[$x])){
        $grid[$x][$y] = 1;
    } else {
        $grid[$x][$y] ++;
    }
}

function addLine(&$grid,$x1,$y1,$x2,$y2){
    if(isHorizontal($x1,$y1,$x2,$y2)){
        for($i = min($y1,$y2); $i <= max($y1,$y2); $i++){
            addPoint($grid,$x1,$i);
        }
    }elseif (isVertical($x1,$y1,$x2,$y2)){
        for($i = min($x1,$x2); $i <= max($x1,$x2); $i++){
            addPoint($grid,$i,$y1);
        }
    }else{
        #diagonal
        #print("LINE:".$x1.":".$y1.",".$x2.":".$y2."\n");
        if($x1 < $x2){
            # fist point is closer to 0
            if($y1 < $y2){
                #increasing
                for($i = $x1; $i <= $x2; $i++){
                    addPoint($grid,$i,$y1+($i-$x1));
                }
            }else{
                #decreasing
                for($i = $x1; $i <= $x2; $i++){
                    addPoint($grid,$i,$y1-($i-$x1));
                }
            }
        }else {
            # second point is closer to 0
            if($y2 < $y1){
                #increasing
                for($i = $x2; $i <= $x1; $i++){
                    addPoint($grid,$i,$y2+($i-$x2));
                }
            }else{
                #decreasing
                for($i = $x2; $i <= $x1; $i++){
                    addPoint($grid,$i,$y2-($i-$x2));
                }
            }
        }
    }
}

if ($handle) {
    $grid = [];
    while (($line = fgets($handle)) !== false) {
        $coords = explode( ',', str_replace(" -> ",",",trim($line) ));
        addLine($grid,$coords[0],$coords[1],$coords[2],$coords[3]);
    }
    fclose($handle);

    $count = 0;
    foreach ($grid as $xkey => $y) {
        foreach ($y as $ykey => $value) {
            if($value > 1){
                $count ++;
            }
        }            
    }
    print($count);
}
?>