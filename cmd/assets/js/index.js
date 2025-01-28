$('.header').on('click', '.search-toggle', function(e) {
  var selector = $(this).data('selector');

  $(selector).toggleClass('show').find('.search-input').focus();
  $(this).toggleClass('active');
  $("#hide_search").toggleClass("active");

  e.preventDefault();
});

$('#hide_search').click(function(e) {
  $('header').removeClass('show');
  $(".search-toggle").removeClass('active');
  $("#hide_search").removeClass("active");
  // e.preventDefault();
});


/*function search(content) {
  var arr = $(".main-content").get();
  for (let i = 0; i < arr.length; i++) {
    if (~arr[i].textContent.indexOf(content)) {
      arr[i].classList.remove("hide");
      continue 
    }
    arr[i].classList.add("hide");
  }
}

const startTime = performance.now()

search("article");   // <---- measured code goes between startTime and endTime

const endTime = performance.now()*/

