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

var showSpinner = function(){
    document.getElementById("spinner").style.display = "block";
}