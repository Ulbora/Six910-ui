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


// $('#myTab a').on('click', function (e) {
//     e.preventDefault()
//     $(this).tab('show')
// })

// $('#myTab prod[href="#prod"]').tab('show')
// $('#myTab cat[href="#cat"]').tab('show')
// // $('#prod-tab').tab('show');
// $('#cat-tab').tab('show');
