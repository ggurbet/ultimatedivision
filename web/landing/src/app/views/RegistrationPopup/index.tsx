// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import React, { useState } from 'react';

import { ChangePassword }
    from '@components/RegistrationPopupPage/ChangePassword';
import { SignUp }
    from '@components/RegistrationPopupPage/SignUp';
import { SignIn }
    from '@components/RegistrationPopupPage/SignIn';

import './index.scss';

export const RegistrationPopup: React.FC<{ handlePopUp: any }> = ({
    handlePopUp
}) => {
    /** check if pop-up components is visible */
    const [isShowPopUpSignIn, setShowPopUpSignIn] = useState(true);
    const [isShowPopUpSignUp, setShowPopUpSignUp] = useState(false);
    const [
        isShowPopUpResetPassword,
        setShowPopUpResetPassword
    ] = useState(false);
    /** show SignInPopUp component */
    const handleSignIn = () => {
        setShowPopUpResetPassword(!isShowPopUpResetPassword);
        setShowPopUpSignIn(!isShowPopUpSignIn);
    };
    /** show SignUpOPopUp component */
    const handleSignUp = () => {
        setShowPopUpSignIn(!isShowPopUpSignIn);
        setShowPopUpSignUp(!isShowPopUpSignUp);
    };
    /** show ResetPasswordPopUp component */
    const handleResetPassword = () => {
        setShowPopUpSignIn(!isShowPopUpSignIn);
        setShowPopUpResetPassword(!isShowPopUpResetPassword);
    };

    return <div className="pop-up-registration">
        <div className="pop-up-registration__wrapper">
            <div
                className="pop-up-registration__wrapper__close"
                onClick={() => handlePopUp()}
            >
                &#x2613;
            </div>
            {isShowPopUpSignIn && <SignIn
                handleResetPassword={handleResetPassword}
                handleSignUp={handleSignUp}
            />}
            {isShowPopUpSignUp && <SignUp handleSignUp={handleSignUp} />}
            {isShowPopUpResetPassword && <ChangePassword
                handleSignIn={handleSignIn}
            />}
        </div>
    </div>;
};
