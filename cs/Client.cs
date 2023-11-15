using System.Net.Http;
using Cosmcs.Client;
using Grpc.Net.Client;
using Grpc.Net.Client.Web;

namespace DecentralCardGame.Cardchain
{
    public class Client
    {
        public EasyClient Ec { get; }
        
        public DecentralCardGame.Cardchain.Cardchain.MsgClient CardchainTxClient { get; }
        public DecentralCardGame.Cardchain.Featureflag.MsgClient FeatureflagTxClient { get; }
        
        public DecentralCardGame.Cardchain.Cardchain.Query.QueryClient CardchainQueryClient { get; }
        public DecentralCardGame.Cardchain.Featureflag.Query.QueryClient FeatureflagQueryClient { get; }
        

        public Client(string rpcUrl, string chainId, byte[] bytes, EasyClientOptions? options = null)
        {
            Ec = new EasyClient(rpcUrl, chainId, bytes, "cosmos", options);
            CardchainTxClient = new DecentralCardGame.Cardchain.Cardchain.MsgClient(Ec);
            FeatureflagTxClient = new DecentralCardGame.Cardchain.Featureflag.MsgClient(Ec);
            
            CardchainQueryClient = new DecentralCardGame.Cardchain.Cardchain.Query.QueryClient(Ec.Channel);
            FeatureflagQueryClient = new DecentralCardGame.Cardchain.Featureflag.Query.QueryClient(Ec.Channel);
            
        }

        public static Client ForUnity(string rpcUrl, string chainId, byte[] bytes)
        {
            return new Client(rpcUrl, chainId, bytes, new EasyClientOptions
            {
                GrpcChannelOptions = new GrpcChannelOptions
                {
                    HttpHandler = new GrpcWebHandler(GrpcWebMode.GrpcWeb, new HttpClientHandler())
                }
            });
        }
    }
}