// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect } from "react";
import { useLocation } from "react-router"

import { UserClient } from "@/api/user";
import { UserService } from "@/user/service";

export const useAuth = () => {
    useEffect(() => {
        checkToken();
    }, []);

    const useQuery = () => {
        return new URLSearchParams(useLocation().search);
    };
    const query = useQuery();

    const userClient = new UserClient();
    const users = new UserService(userClient);

    async function checkToken() {
        return await users.checkToken(query.get('token'));
    };
};
