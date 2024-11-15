using WebSocket = WebSocketSharp.WebSocket;

namespace wsst.Client;

class Program
{
    static void Main(string[] args)
    {
        
        WebSocket wss = new("ws://localhost:8080/ws");
        wss.Connect();
        bool isDead = false;
        
        while (!isDead)
        {
            Console.Write("Send to server:");
            string request = Console.ReadLine()!;

            switch (request)
            {
                case ":x":
                    isDead = true;
                    break;
                case ":init":
                    wss.Send(File.ReadAllBytes(AppDomain.CurrentDomain.BaseDirectory + "init.json"));
                    break;
                case ":score":
                    wss.Send(File.ReadAllBytes(AppDomain.CurrentDomain.BaseDirectory + "score.json"));
                    break;
                case ":text":
                    string data = Console.ReadLine()!;
                    wss.Send(data);
                    break;
                default:
                    Console.WriteLine("Unknown command");
                    break;
            }
        }
        
        Console.ReadLine();
        wss.Close();
    }
}