// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SignIn } from '@/app/components/common/Registration/SignIn';

import closeButton from '@static/img/login/close-icon.svg';

import './index.scss';

// TODO: it will be reworked on wrapper with children props.
export const RegistrationPopup: React.FC<{ closeRegistrationPopup: () => void }> = ({ closeRegistrationPopup }) =>
    <div className="pop-up-registration">
        <div className="pop-up-registration__wrapper">
            <div className="pop-up-registration__wrapper__close" onClick={closeRegistrationPopup}>
                <img src={closeButton} alt="close button" />
            </div>
            <SignIn />
        </div>
    </div>;

