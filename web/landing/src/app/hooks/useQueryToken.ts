// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useLocation } from "react-router";

export const useQueryToken = () => {
    const useQuery = () => {
        return new URLSearchParams(useLocation().search);
    };
    const query = useQuery();
    const token = query.get('token');

    return token;
};
