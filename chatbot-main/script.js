var audio = new Audio('assets/sentmessage.mp3');
var isanewRecording = false;
var mediaRecorder;
var stream = null;
var chunks = [];
var retryCount = 0;
var maxRetries = 3;


function startFunction() {
    setLastSeen();
    introText()
}

function setLastSeen() {
    var date = new Date();
    var lastSeen = document.getElementById("lastseen");
    lastSeen.innerText = "last seen today at " + date.getHours() + ":" + date.getMinutes()
}


function closeFullDP() {
    var x = document.getElementById("fullScreenDP");
    if (x.style.display === 'flex') {
        x.style.display = 'none';
    } else {
        x.style.display = 'flex';
    }
}

function openFullScreenDP() {
    var x = document.getElementById("fullScreenDP");
    if (x.style.display === 'flex') {
        x.style.display = 'none';
    } else {
        x.style.display = 'flex';
    }
}


function isEnter(event) {
    if (event.keyCode == 13) {
        sendMsg();
    }
}


function api_To_Text(input) {

var xhr = new XMLHttpRequest();


xhr.open("POST", "http://localhost:8080/gpt_to_text", true);
xhr.setRequestHeader("Content-Type", "application/json");
xhr.onreadystatechange = function() {
  if (this.readyState == 4) {
    if (this.status == 200) {
     var response = JSON.parse(this.responseText);
     responseSum = response.result
      sendTextMessage(response.result);
    } else if (retryCount < maxRetries) {
      retryCount++;
      setTimeout(api_To_Text(), 1000);
    } else {
        sendTextMessage("Sorry, the server is not available at the moment. Please make sure that your back-end server is up and running.");
    }
  }
};
var data = JSON.stringify({text: input});
xhr.send(data);
}

function api_To_Image(input) {
  return new Promise(function(resolve, reject) {
    var xhr = new XMLHttpRequest();
  
    xhr.open("POST", "http://localhost:8080/gpt_image_gen", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.onreadystatechange = function() {
      if (this.readyState == 4) {
        if (this.status == 200) {
          var response = JSON.parse(this.responseText);
          responseSum = response.result
          resolve(response.result);
        } else if (retryCount < maxRetries) {
          retryCount++;
          setTimeout(api_To_Text(), 1000);
        } else {
          reject("Sorry, the server is not available at the moment. Please make sure that your back-end server is up and running.");
        }
      }
    };
    var data = JSON.stringify({text: input});
    xhr.send(data);
  });
}

async function sendImage() {
  var input = document.getElementById("inputMSG");
  var ti = input.value;
  var lastSeen = document.getElementById("lastseen");
  lastSeen.innerText = "typing...";
  try {
    await api_To_Image(ti);
    input.value = "";
    var textToSend = "/Users/Nonso/Documents/MicroServices/Chat-GPT/example.png";
    setTimeout(setLastSeen, 1000);
    var date = new Date();
    var myLI = document.createElement("li");
    var myDiv = document.createElement("div");
    var greendiv = document.createElement("div");
    var dateLabel = document.createElement("label");
    dateLabel.setAttribute("id", "sentlabel");
    dateLabel.id = "sentlabel";
    dateLabel.innerText = date.getHours() + ":" + date.getMinutes();
    myDiv.setAttribute("class", "received");
    greendiv.setAttribute("class", "grey");

    // Create an image element and set its src attribute to the textToSend parameter
    var imgElement = document.createElement('img');
    imgElement.src = textToSend;

    // Append the image element to the greendiv element
    greendiv.appendChild(imgElement);

    myDiv.appendChild(greendiv);
    myLI.appendChild(myDiv);
    greendiv.appendChild(dateLabel);
    document.getElementById("listUL").appendChild(myLI);
    var s = document.getElementById("chatting");
    s.scrollTop = s.scrollHeight;
    sendTextMessage("Your request to create " + "`" + ti  + "`" + " was created successfully.")
    playSound();
    // updateImageSrc();
  } catch (error) {
    sendTextMessage(error);
  }
}

// function updateImageSrc() {
//   var imgElement = document.querySelector('img[src="/Users/Nonso/Documents/MicroServices/Chat-GPT/example.png"]');
//   if (imgElement) {
//     var newSrc = "/Users/Nonso/Documents/MicroServices/Chat-GPT/new-example.png";
//     imgElement.src = newSrc;
//   }
// }


function api_To_Voice(input) {

}

function sendMsg() {
    var input = document.getElementById("inputMSG");
    var ti = input.value;
    if (input.value == "") {
        return;
    }
    var date = new Date();
    var myLI = document.createElement("li");
    var myDiv = document.createElement("div");
    var greendiv = document.createElement("div");
    var dateLabel = document.createElement("label");
    dateLabel.innerText = date.getHours() + ":" + date.getMinutes();
    myDiv.setAttribute("class", "sent");
    greendiv.setAttribute("class", "green");
    dateLabel.setAttribute("class", "dateLabel");
    greendiv.innerText = input.value;
    myDiv.appendChild(greendiv);
    myLI.appendChild(myDiv);
    greendiv.appendChild(dateLabel);
    document.getElementById("listUL").appendChild(myLI);
    var s = document.getElementById("chatting");
    s.scrollTop = s.scrollHeight;
    input.value = "";
    playSound();
    var lastSeen = document.getElementById("lastseen");
    lastSeen.innerText = "typing...";
    setTimeout(function () { api_To_Text(ti) }, 1500);
    summaryText(ti);
}

function introText() {
    var lastSeen = document.getElementById("lastseen");
    lastSeen.innerText = "typing...";
    var s = document.getElementById("chatting");
    s = "who are you and what can you do to render help to me?"
    api_To_Text(s)
}

function clearChat() {
    document.getElementById("listUL").innerHTML = "";
    waitAndResponce('intro');
}

function sendTextMessage(textToSend) {
    setTimeout(setLastSeen, 1000);
    var date = new Date();
    var myLI = document.createElement("li");
    var myDiv = document.createElement("div");
    var greendiv = document.createElement("div");
    var dateLabel = document.createElement("label");
    dateLabel.setAttribute("id", "sentlabel");
    dateLabel.id = "sentlabel";
    dateLabel.innerText = date.getHours() + ":" + date.getMinutes();
    myDiv.setAttribute("class", "received");
    greendiv.setAttribute("class", "grey");
    greendiv.innerHTML = textToSend;
    myDiv.appendChild(greendiv);
    myLI.appendChild(myDiv);
    greendiv.appendChild(dateLabel);
    document.getElementById("listUL").appendChild(myLI);
    var s = document.getElementById("chatting");
    s.scrollTop = s.scrollHeight;
    playSound();
}


function sendResponse() {
    setTimeout(setLastSeen, 1000);
    var date = new Date();
    var myLI = document.createElement("li");
    var myDiv = document.createElement("div");
    var greendiv = document.createElement("div");
    var dateLabel = document.createElement("label");
    dateLabel.innerText = date.getHours() + ":" + date.getMinutes();
    myDiv.setAttribute("class", "received");
    greendiv.setAttribute("class", "grey");
    dateLabel.setAttribute("class", "dateLabel");
    greendiv.innerText = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum. ";
    myDiv.appendChild(greendiv);
    myLI.appendChild(myDiv);
    greendiv.appendChild(dateLabel);
    document.getElementById("listUL").appendChild(myLI);
    var s = document.getElementById("chatting");
    s.scrollTop = s.scrollHeight;
    playSound();
}

function playSound() {
    audio.play();
}

let buttonClicked = false;

function summaryText(input) {

  if (!buttonClicked) {
    var summeryData = document.getElementById("chatSummery");
    var s = "write a summary of the below text as a question not more than 6 words: /n" + input; 
    var xhr = new XMLHttpRequest();

    xhr.open("POST", "http://localhost:8080/gpt_to_text", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    var data = JSON.stringify({text: s});
    xhr.send(data);

    xhr.onload = function() {
      if (xhr.status === 200) {
        var response = JSON.parse(xhr.responseText);
        if(response.result != "Back-End server not connected to the internet. Please make sure that you're connected to the internet."){
        summeryData.innerText = response.result;
        }
        buttonClicked = true;
      } else {
        console.log("Request failed. Returned status of " + xhr.status);
      }
    }
  } 
}
const myButton = document.getElementById("clickedd");

myButton.addEventListener("click", summaryText());

  
function startRecording() {
  // Get access to the user's microphone or camera
  navigator.mediaDevices.getUserMedia({ audio: true, video: false})
    .then((mediaStream) => {
      // Store the stream for later use
      stream = mediaStream;

      // Create a new MediaRecorder instance
      mediaRecorder = new MediaRecorder(stream);

      // Add event listeners to the MediaRecorder instance
      mediaRecorder.addEventListener('dataavailable', (event) => {
        chunks.push(event.data);
      });

      mediaRecorder.addEventListener('stop', () => {
        // Combine all the recorded chunks into a single Blob object
        let blob = new Blob(chunks, { type: mediaRecorder.mimeType });

        // Create a new URL object from the Blob object
        let url = URL.createObjectURL(blob);

        // Automatically download the recorded media when the user clicks on a link
        let downloadLink = document.createElement('a');
        downloadLink.href = url;
        downloadLink.download = 'audio.webm';
        document.body.appendChild(downloadLink);
        downloadLink.click();

        // Clean up resources
        URL.revokeObjectURL(url);
        chunks = [];
        mediaRecorder = null;
        stream.getTracks().forEach(track => track.stop());
      });

      // Start recording
      mediaRecorder.start();

    })
    .catch((error) => {
      console.error('Error accessing media devices:', error);
    });
}

function stopRecording() {
    // Stop recording
    mediaRecorder.stop();
    console.log("recording ended")
  }

function toggleRecording() {

    if (!isanewRecording) {
      startRecording();
      isanewRecording = true;
    } else {
      stopRecording();
      isanewRecording = false;
    }
  }
