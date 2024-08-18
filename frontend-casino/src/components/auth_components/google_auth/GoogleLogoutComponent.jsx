import { GoogleLogout } from 'react-google-login';

const clientID = "609582150163-ejcmmse6ut85n5iv2sm6s7k4nauirlk8.apps.googleusercontent.com"

function LogoutGo() {
    return (
        <div id="signOutButton">
            <GoogleLogout
                clientId={clientID}
                buttonText={"Logout"}
                onLogoutSuccess={onSuccess}
            />
                
        </div>
    );
}


export default LogoutGo;