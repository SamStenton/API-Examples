import com.google.gson.Gson;

import java.io.*;
import java.net.URL;
import java.util.Arrays;

public class Main {
    /**
     * Print each application string to the console
     * @param apps An array of Strings, the names of the applications
     */
    private static void printApps(String[] apps) {
        Arrays.stream(apps).iterator().forEachRemaining(System.out::println);
    }

    /**
     * Parse response data, return a String[]
     * @param response InputStream of JSON array data, an array of strings, e.g. ["a", "b", "c"]
     * @return Parsed data, as a String[]
     */
    private static String[] parse(InputStream response) {
        return new Gson().fromJson(new InputStreamReader(response), String[].class);
    }

    /**
     * GET request on a given URL
     * @param url The URL to send the GET request to
     * @return InputStream of response data
     * @throws IOException May be thrown by malformed URL, or a 4XX/5XX HTTP error
     */
    private static InputStream getResponse(String url) throws IOException {
        return new URL(url).openStream();
    }

    public static void main(String[] args) throws IOException {
        // Open URL, create InputStream
        InputStream response = getResponse("https://engineering-task.elancoapps.com/api/applications");

        // Parse response to String[]
        String[] parsedResponse = parse(response);

        // Print
        printApps(parsedResponse);
    }
}
