var objects
var first
var last
var sort = false

$( document ).ready(loadData());

$('#customSwitch').change(function() {
  if ((this).checked){
    sort = true
  }else{
    sort = false
  }       
})
 
function loadData() {
  $.get("/products",
  {"sort":sort},
  function(data,status){
    objects = $.parseJSON(data);
    tbody = document.getElementById("table-content");
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
  });
}

function getFirstAndLast(){
  tempdata = document.querySelectorAll('tbody > tr')
  first = tempdata[0].firstElementChild.innerText
  last = tempdata[tempdata.length-1].firstElementChild.innerText
}