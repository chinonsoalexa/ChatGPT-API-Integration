var audio = new Audio('assets/sentmessage.mp3');
var contactString = "<div class='social'> <a target='_blank' href='tel:+916363549133'> <div class='socialItem' id='call'><img class='socialItemI' src='images/phone.svg'/><label class='number'></label></label></div> </a> <a href='mailto:varshithvh@gmail.com'> <div class='socialItem'><img class='socialItemI' src='images/gmail.svg' alt=''></div> </a> <a target='_blank' href='https://github.com/Varshithvhegde'> <div class='socialItem'><img class='socialItemI' src='images/github.svg' alt=''></div> </a> <a target='_blank' href='https://wa.me/916363549133'> <div class='socialItem'><img class='socialItemI' src='images/whatsapp.svg' alt=''>";
var resumeString = "<img src='images/resume_thumbnail.png' class='resumeThumbnail'><div class='downloadSpace'><div class='pdfname'><img src='images/pdf.png'><label>Varshith V Hegde Resume.pdf</label></div><a href='assets/varshith_v_hegde_resume.pdf' download='varshith_v_hegde_resume.pdf'><img class='download' src='images/downloadIcon.svg'></a></div>";
var addressString = "<div class='mapview'><iframe src='https://www.google.com/maps/dir//Moodbidri+private+Bus+Stand,+Bus+Stand+Rd,+Mudbidri,+Karnataka+574227/@13.0639,74.9991985,15z/data=!4m8!4m7!1m0!1m5!1m1!1s0x3ba4ab3d49331379:0x17be05cb5b69caa2!2m2!1d74.9957298!2d13.0680955?hl=en' class='map'></iframe></div><label class='add'><address>B2 'Asara'<br>Kodoli<br>Kolhapur, Maharashtra, INDIA 416114</address>";
var isanewRecording = false;
var mediaRecorder;
var stream = null;
var chunks = [];


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


var retryCount = 0;
var maxRetries = 3;
var responseSum;

function sendData(input) {

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
      setTimeout(sendData(), 1000);
    } else {
        sendTextMessage("Sorry, the server is not available at the moment. Please make sure that your back-end server is up and running.");
    }
  }
};
var data = JSON.stringify({text: input});
xhr.send(data);
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
    setTimeout(function () { sendData(ti) }, 1500);
    summaryText(ti);
}

function introText() {
    var lastSeen = document.getElementById("lastseen");
    lastSeen.innerText = "typing...";
    var s = document.getElementById("chatting");
    s = "who are you and what can you do to render help to me?"
    sendData(s)
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
