document.addEventListener('DOMContentLoaded', function() {
  let vol = document.getElementById('vol');
  let closet = document.getElementById('fishing');
  let playing = false;
  let select = document.getElementById('Element');
  // sortSelect(select);
  select.addEventListener('change', function() {
    console.log(select.options)
    if (!playing){
      closet.setAttribute('src',`./asset/music/${this.value}`)
    } else {
      closet.pause();
      closet.setAttribute('src',`./asset/music/${this.value}`)
      closet.load();
      closet.play();
    }
    
  })

  vol.addEventListener('click', function() {
    if (playing) {
      closet.pause();
      playing = false;
    } else {
      closet.play();
      playing = true;
    }
  });
});
function sortSelect(selElem) {
  let tmpAry = new Array();
  for (var i=0;i<selElem.options.length;i++) {
      tmpAry[i] = new Array();
      tmpAry[i][0] = selElem.options[i].text;
      tmpAry[i][1] = selElem.options[i].value;
  }
  tmpAry.sort();
  while (selElem.options.length > 0) {
      selElem.options[0] = null;
  }
  for (var i=0;i<tmpAry.length;i++) {
      var op = new Option(tmpAry[i][0], tmpAry[i][1]);
      selElem.options[i] = op;
  }
  return;
}