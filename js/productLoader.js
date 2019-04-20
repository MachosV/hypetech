var objects
var sort = false
var direction = -1
var page = 0;
var offset_begin = ""
var offset_end = ""
var pivot = "id"
var maxid
var minid
var first
var last


$ (document ).ready(sortDetector());
$( document ).ready(loadData());
$( document ).ready(prevDetector());
$( document ).ready(nextDetector());


function loadData() {
  $.get("/products",{
    "sort": sort,
    "offset_begin": offset_begin, 
    "offset_end": offset_end,
    "direction": direction,
    "pivot": pivot
  },
  function(data,status){
    objects = $.parseJSON(data);
    tbody = document.getElementById("table-content");
    tbody.innerHTML = ""
    for (var key in objects.products){
      tr = document.createElement("tr");
      for (var field in objects.products[key]){
        td = document.createElement("td");
        td.innerHTML = objects.products[key][field];
        tr.appendChild(td);
      }
      tbody.appendChild(tr);
    }
    getFirstAndLast()
    maxid = objects.metadata.maxid
    minid = objects.metadata.minid
    console.log(minid,maxid)
  });
}

function sortDetector(){
  $('#customSwitch').change(function() {
    if ((this).checked){
      sort = true
      pivot = "pserial"
    }else{
      sort = false
      pivot = "id"
    }
    loadData()       
  })
}

function getFirstAndLast(){
  tempdata = document.querySelectorAll('tbody > tr')
  first = tempdata[0].firstElementChild.innerText
  last = tempdata[tempdata.length-1].firstElementChild.innerText
}

function prevDetector(){
  $('#prev').click(function(){
    direction = 0
    if (page === 0){
      console.log("cannot further");
      return;
    }else{
      page --;
      if (sort){
        offset_begin = first
        offset_end = last
      }else{
        offset_begin = minid
        offset_end = maxid
      }
      loadData()
    } 
  }); 
}

function nextDetector(){
  $('#next').click(function(){
    direction = 1
    page ++;
    if(sort){
      offset_begin = first
      offset_end = last
    } else{
      offset_begin = minid
      offset_end = maxid
    }
    loadData()
  }); 
}