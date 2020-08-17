$(function(){
    $('#userInp').on('input', function(){
        $("#inpText").text("Inputted Text: " + $("#userInp").val());
        $.ajax({
          cache: false,
          type: 'POST',
          async: true,
          url: '/getHash',
          data: {userInp: $("#userInp").val()},
          success: function(data) {
            console.log(data);
            $("#hashText").text("Hashed Text: " + data.hash);
          }
        });
    });
});
