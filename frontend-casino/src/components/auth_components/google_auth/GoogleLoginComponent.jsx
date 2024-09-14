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
                "http://localhost:8081/google/v2/auth/callback/",
                { 
                    googleId: res.profileObj.googleId,
                    email: res.profileObj.email,
                    name: res.profileObj.name,
                    givenName: res.profileObj.givenName,
                    familyName: res.profileObj.familyName,
                    imageUrl: res.profileObj.imageUrl
                },
                {
                    headers: {
                        'Content-Type': 'application/json',
                    },
                }
            );
    
            const token = response.data.token;
            localStorage.setItem("token", token);
            window.location.href = "/";
    
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
        />
    );
}

export default LoginGo;
