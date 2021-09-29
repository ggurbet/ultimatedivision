// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.
import CardCenter from '@static/images/metaverse/card-center.png';
import CardCenterMedium from '@static/images/metaverse/card-center-834.png';
import CardCenterSmall from '@static/images/metaverse/card-center-414.png';
import CardLeft from '@static/images/metaverse/card-left.png';
import CardLeftMedium from '@static/images/metaverse/card-left-834.png';
import CardLeftSmall from '@static/images/metaverse/card-left-414.png';
import CardRight from '@static/images/metaverse/card-right.png';
import CardRightMedium from '@static/images/metaverse/card-right-834.png';
import CardRightSmall from '@static/images/metaverse/card-right-414.png';

import './index.scss';
import { MintButton } from '@components/common/MintButton';

export const Metaverse: React.FC = () => {
    return (
        <section className="metaverse" id="metaverse">
            <div className="wrapper">
                <span className="metaverse__title"
                        data-aos="fade-right"
                        data-aos-duration="600"
                        data-aos-easing="ease-in-out-cubic"
                    >
                        Ultimate Divison
                    </span>
                    <span className="metaverse__subtitle" 
                        data-aos="fade-right"
                        data-aos-duration="600"
                        data-aos-easing="ease-in-out-cubic"
                    >
                        Football Metaverse
                    </span>
                    <div className="metaverse__cards" 
                        data-aos="fade-right"
                        data-aos-duration="600"
                        data-aos-easing="ease-in-out-cubic"
                    >
                    <picture className="left-card">
                        <source media="(max-width: 600px)" srcSet={CardLeftSmall} />
                        <source media="(max-width: 800px)" srcSet={CardLeftMedium} />
                        <source media="(min-width: 1440px)" srcSet={CardLeft}/>
                        <img src={CardLeft} alt="Left player"></img>
                    </picture>
                    <picture className="center-card">
                        <source media="(max-width: 600px)" srcSet={CardCenterSmall} />
                        <source media="(max-width: 800px)" srcSet={CardCenterMedium} />
                        <source media="(min-width: 1440px)" srcSet={CardCenter}/>
                        <img src={CardCenter} alt="Main player"></img>
                    </picture>
                    <picture className="right-card">
                        <source media="(max-width: 600px)" srcSet={CardRightSmall} />
                        <source media="(max-width: 800px)" srcSet={CardRightMedium} />
                        <source media="(min-width: 1440px)" srcSet={CardRight}/>    
                        <img src={CardRight} alt="Right player"></img>
                    </picture>
                </div>
                <div className="metaverse__sold-scale" 
                    data-aos="fade-left"
                    data-aos-duration="600"
                    data-aos-easing="ease-in-out-cubic"
                >
                    <span className="metaverse__sold-scale-text">Cards Sold 0/10000</span>
                </div>
                <MintButton text="MINT"/>
            </div>
        </section>
    );
};
