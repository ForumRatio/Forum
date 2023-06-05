$(document).ready(function() {
    let cat = $('.subtitle');
    let vol = $('#vol');
    let closet = $('#fishing');
    let playing = false;
  
    function Sub(i) {
      loadContent(window.location.hash);
  
      $(document).on('click', 'a', function(event) {
        event.preventDefault();
        var target = $(this).attr('href');
        console.log(target)
        loadContent(target);
        history.pushState(null, null, target);
      });
  
      $(window).on('popstate', function() {
        loadContent(window.location.hash);
      });
  
      function loadContent(target) {
        var url = target;
  
        $.ajax({
          url: url,
          success: function(data) {
            
          }
        });
      }
    }
  
    // cat.eq(0).click(function() {
    //   Sub(1);
    // });
  
    // cat.eq(1).click(function() {
    //   Sub(2);
    // });
  
    // cat.eq(2).click(function() {
    //   Sub(3);
    // });
  
    vol.click(function() {
      if (playing) {
        closet[0].pause();
        playing = false;
      } else {
        closet[0].play();
        playing = true;
      }
    });
  });
  