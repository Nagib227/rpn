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