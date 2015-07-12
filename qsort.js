// quicksort in javascript
//
// A non optimal implementation of quicksort in javascript
// made only for comparative purposes of this
// project.


var quicksort = function( arr ) {

  if(arr.length == 0 ) return arr
  var p = arr.pop()
  var lesser  = arr.filter( function( u ){ return u<p  } )
  var greater = arr.filter( function( u ){ return u>=p } )
  
  return quicksort( lesser ).concat( p, quicksort( greater ) )
}



var arr = [5,9,2,10,12,3,8]

console.log( quicksort(arr) )
