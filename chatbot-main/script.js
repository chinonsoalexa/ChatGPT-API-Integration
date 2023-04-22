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
  return new Promise(function(resolve, reject) {
    var xhr1 = new XMLHttpRequest();
    var xhr2 = new XMLHttpRequest();
    
    // First request
    xhr1.open("POST", "http://localhost:8080/gpt_speach_to_text", true);
    xhr1.setRequestHeader("Content-Type", "application/json");
    xhr1.onreadystatechange = function() {
      if (this.readyState == 4) {
        if (this.status == 200) {
          var response = JSON.parse(this.responseText);
          var speechToTextResult = response.result;
          // Second request
          xhr2.open("POST", "http://localhost:8080/gpt_to_text", true);
          xhr2.setRequestHeader("Content-Type", "application/json");
          xhr2.onreadystatechange = function() {
            if (this.readyState == 4) {
              if (this.status == 200) {
                var response = JSON.parse(this.responseText);
                var textResult = response.result;
                sendTextMessage(textResult);
                resolve(textResult);
              } else {
                reject(new Error("Error making second request"));
              }
            }
          };
          xhr2.send(JSON.stringify({ text: speechToTextResult }));
        } else {
          reject(new Error("Error making first request"));
        }
      }
    };
    xhr1.send(JSON.stringify({ text: input }));
  });
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
    summaryText(ti);
    setTimeout(function () { api_To_Text(ti) }, 1500);
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
    
    api_To_Voice("Users/Nonso/Downloads/audio.webm").then(response => {
      console.log(response); // logs the text response
    }).catch(error => {
      console.error(error); // logs any errors that occur
    });
    
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
