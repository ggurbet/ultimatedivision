// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useDispatch } from 'react-redux';
import { useLocation } from 'react-router';

import { confirmUserEmail } from '@/app/store/actions/users';

/** TODO: Rework this view after design solution */
const ConfirmEmail: React.FC = () => {
    const dispatch = useDispatch();

    const useQuery = () => {
        return new URLSearchParams(useLocation().search);
    };

    const query = useQuery();

    const confirmEmail = () =>
        dispatch(confirmUserEmail(query.get("token")));
    ;

    return <div>
        <input
            value="Confirm Email"
            onClick={confirmEmail}
        />
    </div>;
};

export default ConfirmEmail;
