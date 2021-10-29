// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { lazy } from 'react';
import { Route, Switch } from 'react-router-dom';

const SignIn = lazy(() => import('@/app/views/SignIn'));
const SignUp = lazy(() => import('@/app/views/SignUp'));
const ChangePassword = lazy(() => import('@/app/views/ChangePassword'));
const ConfirmEmail = lazy(() => import('@/app/views/ConfirmEmail'));
const RecoverPassword = lazy(() => import('@/app/views/RecoverPassword'));
const MarketPlace = lazy(() => import('@/app/views/MarketPlacePage'));
const Club = lazy(() => import('@/app/views/ClubPage'));
const Card = lazy(() => import('@/app/views/CardPage'));
const Lot = lazy(() => import('@/app/views/LotPage'));
const FootballField = lazy(() => import('@/app/views/FootballFieldPage'));
const WhitePaper = lazy(() => import('@/app/views/WhitePaperPage'));
const Tokenomics = lazy(() => import('@/app/views/TokenomicsPage'));
const Store = lazy(() => import('@/app/views/StorePage'));

import Summary from '@components/WhitePaper/Summary';
import GameMechanics from '@components/WhitePaper/GameMechanics';
import PayToEarnEconomy from '@components/WhitePaper/PayToEarnEconomy';
import Technology from '@components/WhitePaper/Technology';
import Fund from '@components/Tokenomics/Fund';
import PlayToEarn from '@components/Tokenomics/PlayToEarn';
import Spending from '@components/Tokenomics/Spending';
import Staking from '@components/Tokenomics/Staking';

/** Route base config implementation */
export class ComponentRoutes {
    /** data route config*/
    constructor(
        public path: string,
        public component: React.FC<any>,
        public exact: boolean,
        public children?: ComponentRoutes[]
    ) { }
    /** Method for creating child subroutes path */
    public with(
        child: ComponentRoutes,
        parrent: ComponentRoutes
    ): ComponentRoutes {
        child.path = `${parrent.path}/${child.path}`;

        return this;
    }
    /** Call with method for each child */
    public addChildren(children: ComponentRoutes[]): ComponentRoutes {
        this.children = children.map((item) => item.with(item, this));

        return this;
    }
}

/** Route config implementation */
export class RouteConfig {
    public static SignIn: ComponentRoutes = new ComponentRoutes(
        '/sign-in',
        SignIn,
        true
    );
    public static SignUp: ComponentRoutes = new ComponentRoutes(
        '/sign-up',
        SignUp,
        true
    );
    public static ResetPassword: ComponentRoutes = new ComponentRoutes(
        '/change-password',
        ChangePassword,
        true
    );
    public static ConfirmEmail: ComponentRoutes = new ComponentRoutes(
        '/email/confirm',
        ConfirmEmail,
        true,
    );
    public static RecoverPassword: ComponentRoutes = new ComponentRoutes(
        '/recover-password',
        RecoverPassword,
        true,
    );
    public static MarketPlace: ComponentRoutes = new ComponentRoutes(
        '/marketplace',
        MarketPlace,
        true
    );
    public static Lot: ComponentRoutes = new ComponentRoutes(
        '/lot/:id',
        Lot,
        true
    );
    public static Card: ComponentRoutes = new ComponentRoutes(
        '/card/:id',
        Card,
        false
    );
    public static FootballField: ComponentRoutes = new ComponentRoutes(
        '/field',
        FootballField,
        true
    );
    public static Store: ComponentRoutes = new ComponentRoutes(
        '/store',
        Store,
        true
    );
    public static Club: ComponentRoutes = new ComponentRoutes(
        '/club',
        Club,
        true
    );
    public static Whitepaper: ComponentRoutes = new ComponentRoutes(
        '/whitepaper',
        WhitePaper,
        false
    );
    public static Tokenomics: ComponentRoutes = new ComponentRoutes(
        '/tokenomics',
        Tokenomics,
        false
    );
    public static Summary: ComponentRoutes = new ComponentRoutes(
        'summary',
        Summary,
        true
    );
    public static GameMechanics: ComponentRoutes = new ComponentRoutes(
        'game-mechanics',
        GameMechanics,
        true
    );
    public static PayToEarnEconomy: ComponentRoutes = new ComponentRoutes(
        'pay-to-earn-and-economy',
        PayToEarnEconomy,
        true
    );
    public static Technology: ComponentRoutes = new ComponentRoutes(
        'technology',
        Technology,
        true
    );
    public static Spending: ComponentRoutes = new ComponentRoutes(
        'udt-spending',
        Spending,
        true
    );
    public static PayToEarn: ComponentRoutes = new ComponentRoutes(
        'pay-to-earn',
        PlayToEarn,
        true
    );
    public static Staking: ComponentRoutes = new ComponentRoutes(
        'staking',
        Staking,
        true
    );
    public static Fund: ComponentRoutes = new ComponentRoutes(
        'ud-dao-fund',
        Fund,
        true
    );
    public static Default: ComponentRoutes = new ComponentRoutes(
        '/',
        SignIn,
        true
    );
    public static routes: ComponentRoutes[] = [
        RouteConfig.Default,
        RouteConfig.FootballField,
        RouteConfig.MarketPlace,
        RouteConfig.Club,
        RouteConfig.Card,
        RouteConfig.Lot,
        RouteConfig.Store,
        RouteConfig.SignIn,
        RouteConfig.SignUp,
        RouteConfig.ResetPassword,
        RouteConfig.ConfirmEmail,
        RouteConfig.RecoverPassword,
        RouteConfig.Whitepaper.addChildren([
            RouteConfig.Summary,
            RouteConfig.GameMechanics,
            RouteConfig.PayToEarnEconomy,
            RouteConfig.Technology,
        ]),
        RouteConfig.Tokenomics.addChildren([
            RouteConfig.Spending,
            RouteConfig.PayToEarn,
            RouteConfig.Staking,
            RouteConfig.Fund,
        ]),
    ];
}

export const Routes = () =>
    <Switch>
        {RouteConfig.routes.map((route, index) =>
            <Route
                key={index}
                path={route.path}
                component={route.component}
                exact={route.exact}
            />
        )}
    </Switch>;

