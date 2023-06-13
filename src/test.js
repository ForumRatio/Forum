let inputElement = document.querySelector('.comecrit');
let conversationElement = document.querySelector('.conversation');
let bubbleIdCounter = 1;

function toggleLike(element) {
  let bubbleElement = element.closest('.conv');
  let likeCountElement = bubbleElement.querySelector('.poucebloClics');
  let likeCount = parseInt(likeCountElement.getAttribute('data-count'));

  if (element.classList.contains('active')) {
    // Annuler le like
    likeCount--;
    element.classList.remove('active');
  } else {
    // Ajouter le like
    likeCount++;
    element.classList.add('active');
  }

  likeCountElement.setAttribute('data-count', likeCount);
  likeCountElement.textContent = likeCount;
}

function toggleDislike(element) {
  let bubbleElement = element.closest('.conv');
  let dislikeCountElement = bubbleElement.querySelector('.poucerougeClics');
  let dislikeCount = parseInt(dislikeCountElement.getAttribute('data-count'));

  if (element.classList.contains('active')) {
    // Annuler le dislike
    dislikeCount--;
    element.classList.remove('active');
  } else {
    // Ajouter le dislike
    dislikeCount++;
    element.classList.add('active');
  }

  dislikeCountElement.setAttribute('data-count', dislikeCount);
  dislikeCountElement.textContent = dislikeCount;
}

function toggleFuck(element) {
  let bubbleElement = element.closest('.conv');
  let fuckCountElement = bubbleElement.querySelector('.fuckClics');
  let fuckCount = parseInt(fuckCountElement.getAttribute('data-count'));

  if (element.classList.contains('active')) {
    // Annuler le fuck
    fuckCount--;
    element.classList.remove('active');
  } else {
    // Ajouter le fuck
    fuckCount++;
    element.classList.add('active');
  }

  fuckCountElement.setAttribute('data-count', fuckCount);
  fuckCountElement.textContent = fuckCount;
}

inputElement.addEventListener('keyup', function(event) {
  if (event.keyCode === 13) {
    let message = document.createElement('p');
    message.className = 'message';
    message.innerHTML = inputElement.value;

    let bubbleElement = document.createElement('div');
    let bubbleId = `bubble-${bubbleIdCounter}`;
    bubbleElement.id = bubbleId;
    bubbleElement.classList.add('conv');
    bubbleElement.innerHTML = `
      <img class="pdp" src="/asset/igypdp.png">
      <img class="bulletext" src="/asset/bulles.png">
      <div class="container">
        <div class="blo">
          <span id="poucebloClics-${bubbleIdCounter}" class="poucebloClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="pouceblo" src="/asset/cool.png" onclick="toggleLike(this)"></input>
        </div>
        <div class="rouge">
          <span id="poucerougeClics-${bubbleIdCounter}" class="poucerougeClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="poucerouge" src="/asset/ugh.png" onclick="toggleDislike(this)"></input>
        </div>
        <div class="fucks">
          <span id="fuckClics-${bubbleIdCounter}" class="fuckClics" data-count="0" style="font-size: 200%;">0</span>
          <input type="image" class="fuck" src="/asset/duh.png" onclick="toggleFuck(this)"></input>
        </div>
      </div>
    `;
    
    bubbleIdCounter++;

    bubbleElement.appendChild(message);
    conversationElement.appendChild(bubbleElement);

    inputElement.value = '';
  }
});

