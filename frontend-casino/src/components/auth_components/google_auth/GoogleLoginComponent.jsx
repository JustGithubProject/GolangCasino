import { GoogleLogin } from 'react-google-login';
import axios from 'axios';

const clientID = "609582150163-ejcmmse6ut85n5iv2sm6s7k4nauirlk8.apps.googleusercontent.com"


function LoginGo() {

    const onSuccess = async (res) => {
        console.log("LOGIN SUCCESS! Current user: ", res.profileObj);
        // const response = await axios.post(
        //     "http://127.0.0.1:8081/google/auth/callback/",
        //     res.profileObj,
        //     {
        //         withCredentials: true,
        //         headers: {
        //           'Content-Type': 'application/json',
        //         },
        //     }
        // )
        // const token = response.token;
        const response = await axios.get(
            "http://127.0.0.1:8081/google/oauth/"
        )
        console.log(response.data);

        // localStorage.setItem("google_token", token);
    }

    const onFailure = (res) => {
        console.log("LOGIN FAILED res: ", res);
    }


    return (
        <GoogleLogin
            clientId={clientID}
            onSuccess={onSuccess}
            onFailure={onFailure}
            cookiePolicy={'single_host_origin'}
            isSignedIn={true}
        />
    );
}

export default LoginGo;
