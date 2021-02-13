import React, { useEffect, useRef } from 'react';
import PropTypes from 'prop-types';
import styles from './index.module.css';

const propTypes = {
    isLoading: PropTypes.bool.isRequired,
    isLoaded: PropTypes.bool.isRequired,
    onFetch: PropTypes.func.isRequired,
    children: PropTypes.node.isRequired,
};

function disconnectObserver(observer) {
    if (observer) {
        observer.disconnect();
    }
}

export default function InfiniteScroll({ isLoading, isLoaded, onFetch, children }) {
    const observerRef = useRef(null);
    const intersectionRef = useRef(null);

    useEffect(() => {
        disconnectObserver(observerRef.current);

        observerRef.current = new IntersectionObserver(([node]) => {
            if (node.isIntersecting && !isLoaded && !isLoading) {
                onFetch();
            }
        });

        if (intersectionRef.current) {
            observerRef.current.observe(intersectionRef.current);
        }

        return () => disconnectObserver(observerRef.current);
    }, [isLoading, isLoaded]);

    return (
        <>
            {children}
            <div ref={intersectionRef} className={styles.intersection} />
        </>
    );
}

InfiniteScroll.propTypes = propTypes;
