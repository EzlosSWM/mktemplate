package statics

var Html string = (`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>@EzlosSWM</title>
    <link rel="stylesheet" href="css/style.css">
</head>
<body>
    
    <div class="main-body">
        <h1>MKTEMPLATE</h1>
        <h3>MKTEMPLATE is a command line tool created using Golang with the intent to automate the creation of a simple front-end web project.</h3><br>

        <p>
            Currently, there is a primary index.html file and the css folder containing style.css for custom styling.<br>
            
            If you want to assist in improving the underlying code, please contact me on my social medias:
        </p>
        <div class="socials">
            <ul>
                <li><a href="https://github.com/EzlosSWM" target="_blank"><img src="https://i.imgur.com/GT6NeDr.png" class="img" alt="Twitter"></a></li>
                <li><a href="https://twitter.com/EzlosSWM" target="_blank"><img src="https://i.imgur.com/uVfSOoZ.png" class="img" alt="Github"></a></li>
            </ul>

            
        </div>
        <div class="wavy">
            <span style="--i:1;">E</span>
            <span style="--i:2;">Z</span>
            <span style="--i:3;">L</span>
            <span style="--i:4;">O</span>
            <span style="--i:5;">S</span>
            <span style="--i:6;">-</span>
            <span style="--i:7;">S</span>
            <span style="--i:8;">W</span>
            <span style="--i:9;">M</span>
        </div>
    </div>
    <script src="app/app.js"></script>
</body>
</html>
`)

var CssFile string = (`
.main-body {

    color: #fff;
    text-align: center;
    font-size: large;
}

h1 {

    margin-bottom: 4rem;

}

ul {

    list-style-type: none;

}

a {

    text-decoration: none;
    color: inherit;

}

body {

    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    background: #000;

}

.img {

    width: 5rem;
    height:fit-content;

}


.wavy {

    letter-spacing: 0.5em;
    position: relative;
    -webkit-box-reflect: below -12px linear-gradient(transparent,rgba(0,0,0,0.2));

}

.wavy span {

    position: relative;
    display: inline-block;
    color: #fff;
    font-size: 2rem;
    animation: animate 1s ease-in-out infinite;
    animation-delay: calc(0.1s * var(--i));

}

@keyframes animate {

    0% {
        transform: translateY(0px);
    }
    20% {
        transform: translateY(-20px);
    }
    40%, 100% {
		transform: translateY(0px);
    }

}
`)

var JS string = (`alert("Thank you for downloading mktemplate!")`)
