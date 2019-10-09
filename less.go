package sortedlist


var LessInt= func(i Score,j Score)bool{
	if i.(int)<j.(int){return true} else {return false}
}

var LessInt64= func(i Score,j Score)bool{
	if i.(int64)<j.(int64){return true} else {return false}
}

var LessString= func(i Score,j Score)bool{
	if i.(string)<j.(string){return true} else {return false}
}
