jQuery(document).ready(function ($) {
    $(".clickable-row").click(function () {
        window.location = $(this).data("href");
    });
});


// $('#thumbnail').on('change',function(){
//     alert("test");
//     var fileName = $(this).val();
// })

$(document).ready(function () {
    $("#thumbnail").keyup(function () {
        //alert( $(this).val());
        var imageName = $(this).val();
        document.getElementById("thumbnailImg").src = imageName;
    });
});

$(document).ready(function () {
    $("#image1").keyup(function () {
        //alert( $(this).val());
        var imageName = $(this).val();
        document.getElementById("image1Img").src = imageName;
    });
});



$(document).ready(function () {
    $("#image2").keyup(function () {
        //alert( $(this).val());
        var imageName = $(this).val();
        document.getElementById("image2Img").src = imageName;
    });
});



$(document).ready(function () {
    $("#image3").keyup(function () {
        //alert( $(this).val());
        var imageName = $(this).val();
        document.getElementById("image3Img").src = imageName;
    });
});



$(document).ready(function () {
    $("#image4").keyup(function () {
        //alert( $(this).val());
        var imageName = $(this).val();
        document.getElementById("image4Img").src = imageName;
    });
});


$(document).ready(function () {
    $("#image").keyup(function () {
        //alert( $(this).val());
        var imageName = $(this).val();
        document.getElementById("imageImg").src = imageName;
    });
});


$(document).ready(function () {
    $("#image").load(function () {
        //alert( $(this).val());
        var imageName = $(this).val();
        document.getElementById("imageImg").src = imageName;
    });
});



$(document).ready(function () {
    $("#logo").keyup(function () {
        //alert( $(this).val());
        var imageName = $(this).val();
        document.getElementById("logoImg").src = imageName;
    });
});



// Callback that creates and populates a data table,
  // instantiates the pie chart, passes in the data and
  // draws it.
//   function drawVisitorChart() {

//     // Create the data table.
//     var data = new google.visualization.DataTable();
//     data.addColumn('string', 'Day');
//     data.addColumn('number', 'Visitors');
//     data.addRows([
//       ['11/1', 30],
//       ['11/2', 10],
//       ['11/3', 10],
//       ['11/4', 10],
//       ['11/5', 20]
//     ]);

//     // Set chart options
//     var options = {'title':'Visitor Count by Day',
//                    'width':500,
//                    'height':400};

//     // Instantiate and draw our chart, passing in some options.
//     var chart = new google.visualization.ColumnChart(document.getElementById('visitor_chart_div'));
//     chart.draw(data, options);
//   }

function drawVisitorChart(dataval) {

    // Create the data table.
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Day');
    data.addColumn('number', 'Visitors');
    data.addRows(dataval);

    // Set chart options
    var options = {'title':'Visitor Count by Day',
                   'width':500,
                   'height':400};

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.ColumnChart(document.getElementById('visitor_chart_div'));
    chart.draw(data, options);
  }


// Callback that creates and populates a data table,
  // instantiates the pie chart, passes in the data and
  // draws it.
//   function drawSalesChart() {

//     // Create the data table.
//     var data = new google.visualization.DataTable();
//     data.addColumn('string', 'Day');
//     data.addColumn('number', '$ Sales');
//     data.addRows([
//       ['11/1', 300.25],
//       ['11/2', 100.15],
//       ['11/3', 101.25],
//       ['11/4', 1001.55],
//       ['11/5', 201.55]
//     ]);

//     // Set chart options
//     var options = {'title':'Sales by Day',
//                    'width':500,
//                    'height':400};

//     // Instantiate and draw our chart, passing in some options.
//     var chart = new google.visualization.ColumnChart(document.getElementById('sales_chart_div'));
//     chart.draw(data, options);
//   }

function drawSalesChart(dataval) {

    // Create the data table.
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Day');
    data.addColumn('number', '$ Sales');
    data.addRows(dataval);

    // Set chart options
    var options = {'title':'Sales by Day',
                   'width':500,
                   'height':400};

    // Instantiate and draw our chart, passing in some options.
    var chart = new google.visualization.ColumnChart(document.getElementById('sales_chart_div'));
    chart.draw(data, options);
  }

  var loadVisitors = function(data){
    google.charts.setOnLoadCallback(drawVisitorChart(data));
  }

  var loadSales = function(data){
    google.charts.setOnLoadCallback(drawSalesChart(data));
  }


function checkPasswordMatch() {
    var password = $("#password").val();
    var confirmPassword = $("#password2").val();
    if (password != confirmPassword) {       
        document.getElementById("CheckPasswordMatch").style.visibility = "visible";
        $("#CheckPasswordMatch").html("Passwords does not match!");
    }
    else {
        document.getElementById("CheckPasswordMatch").style.visibility = "hidden";
    }
}

$(document).ready(function () {
    $("#password").keyup(checkPasswordMatch);
});
$(document).ready(function () {
    $("#password2").keyup(checkPasswordMatch);
});