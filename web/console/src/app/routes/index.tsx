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
const UserCards = lazy(() => import('@/app/views/UserCards'));
const Card = lazy(() => import('@/app/views/CardPage'));
const Lot = lazy(() => import('@/app/views/LotPage'));
const Field = lazy(() => import('@/app/views/FieldPage'));
const WhitePaper = lazy(() => import('@/app/views/WhitePaperPage'));
const Tokenomics = lazy(() => import('@/app/views/TokenomicsPage'));
const Store = lazy(() => import('@/app/views/StorePage'));
const Division = lazy(() => import('@/app/views/Division'));
const Match = lazy(() => import('@/app/views/Match'));
const MatchFinder = lazy(() => import('@components/common/MatchFinder'));
const Home = lazy(() => import('@/app/views/Home'));
const Navbar = lazy(() => import('@/app/components/common/Navbar'));
const PlayerProfile = lazy(() => import('@/app/views/PlayerProfile'));

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
        public component: any,
        public exact: boolean,
        public children?: ComponentRoutes[]
    ) {}
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
    public static Division: ComponentRoutes = new ComponentRoutes(
        /** TODO: it will be replaced with id parameter */
        '/divisions',
        Division,
        true
    );
    public static Field: ComponentRoutes = new ComponentRoutes(
        '/field',
        Field,
        true
    );
    public static Store: ComponentRoutes = new ComponentRoutes(
        '/store',
        Store,
        true
    );
    public static Cards: ComponentRoutes = new ComponentRoutes(
        '/cards',
        UserCards,
        true
    );
    public static Match: ComponentRoutes = new ComponentRoutes(
        '/match',
        Match,
        true
    );
    public static Home: ComponentRoutes = new ComponentRoutes(
        '/home',
        Home,
        true
    );
    public static PlayerProfile: ComponentRoutes = new ComponentRoutes(
        '/player-profile',
        PlayerProfile,
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

    public static routes: ComponentRoutes[] = [
        RouteConfig.Home,
        RouteConfig.Field,
        RouteConfig.MarketPlace,
        RouteConfig.Cards,
        RouteConfig.Card,
        RouteConfig.Division,
        RouteConfig.Lot,
        RouteConfig.Store,
        RouteConfig.Match,
        RouteConfig.PlayerProfile,
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

/** Route config that implements auth actions */
export class AuthRouteConfig {
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
    public static ChangePassword: ComponentRoutes = new ComponentRoutes(
        '/change-password',
        ChangePassword,
        true
    );
    public static ConfirmEmail: ComponentRoutes = new ComponentRoutes(
        '/email/confirm',
        ConfirmEmail,
        true
    );
    public static ResetPassword: ComponentRoutes = new ComponentRoutes(
        '/reset-password',
        RecoverPassword,
        true
    );
    public static Default: ComponentRoutes = new ComponentRoutes(
        '/',
        RouteConfig.Home.component,
        true
    );
    public static routes: ComponentRoutes[] = [
        AuthRouteConfig.ConfirmEmail,
        AuthRouteConfig.Default,
        AuthRouteConfig.ResetPassword,
        AuthRouteConfig.ChangePassword,
        AuthRouteConfig.SignIn,
        AuthRouteConfig.SignUp,
    ];
}

export const Routes = () =>
    <Switch>
        {AuthRouteConfig.routes.map((route, index) =>
            <Route
                key={index}
                path={route.path}
                component={route.component}
                exact={route.exact}
            />
        )}
        <Route>
            <Navbar />
            <MatchFinder />
            {RouteConfig.routes.map((route, index) =>
                <Route
                    key={index}
                    path={route.path}
                    component={route.component}
                    exact={route.exact}
                />
            )}
        </Route>
    </Switch>;

