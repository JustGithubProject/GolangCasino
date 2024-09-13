import { GoogleLogin } from 'react-google-login';
import axios from 'axios';

const clientID = "609582150163-ejcmmse6ut85n5iv2sm6s7k4nauirlk8.apps.googleusercontent.com"


function LoginGo() {

    const onSuccess = async (res) => {
        console.log("Current user: ", res.profileObj);
        try {
            const { tokenId } = res;
            console.log("TokenID: ", tokenId);
            const response = await axios.post(
                "http://localhost:8081/google/v2/auth/callback/?provider=google",
                { id_token: tokenId },
                {
                    withCredentials: true,
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }
            );
    
            const { token } = response.data;
            localStorage.setItem("google_token", token);
    
            console.log("LOGIN SUCCESS! Token: ", token);
        } catch (error) {
            console.log("Error during Google login", error);
        }
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
