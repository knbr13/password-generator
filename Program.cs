namespace PasswordGenerator
{
    class Program
    {
        static int GetPasswordLength()
        {
            int passwordLength;
            while (true)
            {
                Console.Write("Enter the length of the password: ");
                if (int.TryParse(Console.ReadLine(), out passwordLength) && passwordLength > 0)
                {
                    break;
                }
                Console.WriteLine("Invalid input. Please enter a positive integer.");
            }
            return passwordLength;
        }

        static bool GetBooleanInput(string message)
        {
            while (true)
            {
                Console.Write(message);
                if (message == "Y" || message == "YES")
                {
                    return true;
                }
                else if (message == "N" || message == "NO")
                {
                    return false;
                }
                Console.WriteLine("Invalid input. Please enter 'Y' or 'N'.");
            }
        }
    }
}
