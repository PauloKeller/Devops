package com.paulokeller;

import junit.framework.Test;
import junit.framework.TestCase;
import junit.framework.TestSuite;

/**
 * Unit test for simple App.
 */
public class AppTest 
    extends TestCase
{
    /**
     * Create the test case
     *
     * @param testName name of the test case
     */
    public AppTest( String testName )
    {
        super( testName );
    }

    /**
     * @return the suite of tests being tested
     */
    public static Test suite()
    {
        return new TestSuite( AppTest.class );
    }

    /**
     * Rigourous Test :-)
     */
    public void testSubtract() {
        App app = new App();

        assertEquals(5,app.subtractNumbers(10,5));
    }

    public void testAdd() {
        App app = new App();

        assertEquals(10,app.addNumbers(5,5));
    }
}
